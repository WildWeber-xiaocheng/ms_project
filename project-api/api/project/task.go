package project

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
	"test.com/project-api/pkg/model"
	"test.com/project-api/pkg/model/pro"
	"test.com/project-api/pkg/model/tasks"
	common "test.com/project-common"
	"test.com/project-common/errs"
	"test.com/project-grpc/task"
	"time"
)

type HandlerTask struct {
}

func (t *HandlerTask) taskStages(c *gin.Context) {
	result := &common.Result{}
	//1. 校验参数，验证参数的合法性
	projectCode := c.PostForm("projectCode")
	page := &model.Page{}
	page.Bind(c)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	//2. 调用grpc服务
	msg := &task.TaskReqMessage{
		MemberId:    c.GetInt64("memberId"),
		ProjectCode: projectCode,
		Page:        page.Page,
		PageSize:    page.PageSize,
	}
	stages, err := TaskServiceClient.TaskStages(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}
	//3. 处理响应
	var resp []*tasks.TaskStagesResp
	copier.Copy(&resp, stages.List)
	if resp == nil {
		resp = []*tasks.TaskStagesResp{}
	}
	for _, v := range resp {
		v.TasksLoading = true  //任务加载状态
		v.FixedCreator = false //添加任务按钮定位
		v.ShowTaskCard = false //是否显示创建卡片
		v.Tasks = []int{}
		v.DoneTasks = []int{}
		v.UnDoneTasks = []int{}
	}

	c.JSON(http.StatusOK, result.Success(gin.H{
		"list":  resp,
		"total": stages.Total,
		"page":  page.Page,
	}))
}

func (t *HandlerTask) memberProjectList(c *gin.Context) {
	result := &common.Result{}
	projectCode := c.PostForm("projectCode")
	page := &model.Page{}
	page.Bind(c)

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Second)
	defer cancel()
	msg := &task.TaskReqMessage{
		MemberId:    c.GetInt64("memberId"),
		ProjectCode: projectCode,
		Page:        page.Page,
		PageSize:    page.PageSize,
	}

	memberResp, err := TaskServiceClient.MemberProjectList(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}
	var resp []*pro.MemberProjectResp
	copier.Copy(&resp, memberResp.List)
	if resp == nil {
		resp = []*pro.MemberProjectResp{}
	}
	c.JSON(http.StatusOK, result.Success(gin.H{
		"list":  resp,
		"total": memberResp.Total,
		"page":  page.Page,
	}))
}

func (t *HandlerTask) taskList(c *gin.Context) {
	result := &common.Result{}
	stageCode := c.PostForm("stageCode")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	list, err := TaskServiceClient.TaskList(ctx, &task.TaskReqMessage{StageCode: stageCode})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}
	var taskDisplayList []*tasks.TaskDisplay
	copier.Copy(&taskDisplayList, list.List)
	if taskDisplayList == nil {
		taskDisplayList = []*tasks.TaskDisplay{}
	}
	for _, v := range taskDisplayList {
		if v.Tags == nil {
			v.Tags = []int{}
		}
		if v.ChildCount == nil {
			v.ChildCount = []int{}
		}
	}
	c.JSON(http.StatusOK, result.Success(taskDisplayList))
}

func NewTask() *HandlerTask {
	return &HandlerTask{}
}
