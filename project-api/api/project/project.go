package project

import (
	"context"
	"github.com/WildWeber-xiaocheng/ms_project/project-api/pkg/model"
	"github.com/WildWeber-xiaocheng/ms_project/project-api/pkg/model/pro"
	common "github.com/WildWeber-xiaocheng/ms_project/project-common"
	"github.com/WildWeber-xiaocheng/ms_project/project-common/errs"
	"github.com/WildWeber-xiaocheng/ms_project/project-grpc/project"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
	"strconv"
	"time"
)

type HandlerProject struct {
}

// index 升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Success 200
// @Router /project/index [get]
func (p HandlerProject) index(c *gin.Context) {
	result := &common.Result{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	msg := &project.IndexMessage{}
	indexResponse, err := ProjectServiceClient.Index(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}
	var menus []*model.Menu
	copier.Copy(&menus, indexResponse.Menus)
	c.JSON(http.StatusOK, result.Success(menus))
}

func (p HandlerProject) myProjectList(c *gin.Context) {
	//1. 获取参数
	result := &common.Result{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	memberId := c.GetInt64("memberId")
	memberName := c.GetString("memberName")
	page := &model.Page{}
	page.Bind(c)
	selectBy := c.PostForm("selectBy")
	msg := &project.ProjectRpcMessage{
		MemberId:   memberId,
		MemberName: memberName,
		SelectBy:   selectBy,
		Page:       page.Page,
		PageSize:   page.PageSize,
	}
	myProjectResponse, err := ProjectServiceClient.FindProjectByMemId(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}
	var pms []*pro.ProjectAndMember
	copier.Copy(&pms, myProjectResponse.Pm)
	if pms == nil {
		pms = []*pro.ProjectAndMember{}
	}
	c.JSON(http.StatusOK, result.Success(gin.H{
		"list":  pms,
		"total": myProjectResponse.Total,
	}))
}

func (p HandlerProject) projectTemplate(c *gin.Context) {
	//1. 获取参数
	result := &common.Result{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	memberId := c.GetInt64("memberId")
	memberName := c.GetString("memberName")
	organizationCode := c.GetString("organizationCode")
	var page = &model.Page{}
	page.Bind(c)
	viewTypeStr := c.PostForm("viewType")
	viewType, _ := strconv.ParseInt(viewTypeStr, 10, 64)

	projectTemplateRsp, err := ProjectServiceClient.FindProjectTemplate(ctx,
		&project.ProjectRpcMessage{
			MemberId:         memberId,
			MemberName:       memberName,
			OrganizationCode: organizationCode,
			Page:             page.Page,
			PageSize:         page.PageSize,
			ViewType:         int32(viewType),
		})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}
	var pts []*pro.ProjectTemplate
	copier.Copy(&pts, projectTemplateRsp.Ptm)
	if pts == nil {
		pts = []*pro.ProjectTemplate{}
	}
	for _, v := range pts {
		if v.TaskStages == nil {
			v.TaskStages = []*pro.TaskStagesOnlyName{}
		}
	}
	c.JSON(http.StatusOK, result.Success(gin.H{
		"list":  pts,
		"total": projectTemplateRsp.Total,
	}))
}

func (p HandlerProject) projectSave(c *gin.Context) {
	result := &common.Result{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	memberId := c.GetInt64("memberId")
	organizationCode := c.GetString("organizationCode")

	var req *pro.SaveProjectRequest
	c.ShouldBind(&req)
	msg := &project.ProjectRpcMessage{
		MemberId:         memberId,
		Name:             req.Name,
		OrganizationCode: organizationCode,
		Description:      req.Description,
		TemplateCode:     req.TemplateCode,
		Id:               int64(req.Id)}
	saveProjectMessage, err := ProjectServiceClient.SaveProject(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}
	var sp *pro.SaveProject
	copier.Copy(&sp, saveProjectMessage)
	c.JSON(http.StatusOK, result.Success(sp))
}

func (p HandlerProject) readProject(c *gin.Context) {
	result := &common.Result{}
	projectCode := c.PostForm("projectCode")
	memberId := c.GetInt64("memberId")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	detail, err := ProjectServiceClient.FindProjectDetail(ctx, &project.ProjectRpcMessage{ProjectCode: projectCode, MemberId: memberId})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}
	pd := &pro.ProjectDetail{}
	copier.Copy(pd, detail)
	c.JSON(http.StatusOK, result.Success(pd))
}

func (p HandlerProject) recycleProject(c *gin.Context) {
	result := &common.Result{}
	projectCode := c.PostForm("projectCode")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err := ProjectServiceClient.UpdateDeletedProject(ctx, &project.ProjectRpcMessage{ProjectCode: projectCode, Deleted: true})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}
	c.JSON(http.StatusOK, result.Success([]int{}))
}

func (p HandlerProject) recoveryProject(c *gin.Context) {
	result := &common.Result{}
	projectCode := c.PostForm("projectCode")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err := ProjectServiceClient.UpdateDeletedProject(ctx, &project.ProjectRpcMessage{ProjectCode: projectCode, Deleted: false})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}
	c.JSON(http.StatusOK, result.Success([]int{}))
}

func (p HandlerProject) collectProject(c *gin.Context) {
	result := &common.Result{}
	projectCode := c.PostForm("projectCode")
	collectType := c.PostForm("type")
	memberId := c.GetInt64("memberId")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err := ProjectServiceClient.UpdateCollectProject(ctx, &project.ProjectRpcMessage{ProjectCode: projectCode, CollectType: collectType, MemberId: memberId})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}
	c.JSON(http.StatusOK, result.Success([]int{}))
}

func (p HandlerProject) editProject(c *gin.Context) {
	result := &common.Result{}
	var req *pro.ProjectReq
	_ = c.ShouldBind(&req)
	memberId := c.GetInt64("memberId")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	msg := &project.UpdateProjectMessage{}
	copier.Copy(msg, req)
	msg.MemberId = memberId
	_, err := ProjectServiceClient.UpdateProject(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}
	c.JSON(http.StatusOK, result.Success([]int{}))
}

func (p HandlerProject) getLogBySelfProject(c *gin.Context) {
	result := &common.Result{}
	var page = &model.Page{}
	page.Bind(c)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	msg := &project.ProjectRpcMessage{
		MemberId: c.GetInt64("memberId"),
		Page:     page.Page,
		PageSize: page.PageSize,
	}
	projectLogResponse, err := ProjectServiceClient.GetLogBySelfProject(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}
	var list []*model.ProjectLog
	copier.Copy(&list, projectLogResponse.List)
	if list == nil {
		list = []*model.ProjectLog{}
	}
	c.JSON(http.StatusOK, result.Success(list))
}

func (p *HandlerProject) nodeList(c *gin.Context) {
	result := &common.Result{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	response, err := ProjectServiceClient.NodeList(ctx, &project.ProjectRpcMessage{})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}
	var list []*model.ProjectNodeTree
	copier.Copy(&list, response.Nodes)
	c.JSON(http.StatusOK, result.Success(gin.H{
		"nodes": list,
	}))
}

func (p *HandlerProject) FindProjectByMemberId(memberId int64, projectCode string, taskCode string) (*pro.Project, bool, bool, *errs.BError) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	msg := &project.ProjectRpcMessage{
		MemberId:    memberId,
		ProjectCode: projectCode,
		TaskCode:    taskCode,
	}
	projectResponse, err := ProjectServiceClient.FindProjectByMemberId(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		return nil, false, false, errs.NewError(errs.ErrorCode(code), msg)
	}
	if projectResponse.Project == nil {
		return nil, false, false, nil
	}
	pr := &pro.Project{}
	copier.Copy(pr, projectResponse.Project)
	return pr, true, projectResponse.IsOwner, nil
}
func New() *HandlerProject {
	return &HandlerProject{}
}
