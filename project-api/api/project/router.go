package project

import (
	"github.com/WildWeber-xiaocheng/ms_project/project-api/api/midd"
	"github.com/WildWeber-xiaocheng/ms_project/project-api/router"
	"github.com/gin-gonic/gin"
	"log"
)

// Route 接口的实现类
type RouterProject struct {
}

func init() {
	log.Println("init project router")
	ru := &RouterProject{}
	router.RegisterR(ru)
}

func (*RouterProject) Register(r *gin.Engine) {
	InitRpcProjectClient() //初始化grpc客户端连接
	h := New()
	group := r.Group("/project")
	group.Use(midd.TokenVerify())
	//group.Use(Auth())
	//group.Use(ProjectAuth())
	group.POST("/index", h.index)
	group.POST("/project/selfList", h.myProjectList)
	group.POST("/project", h.myProjectList) //和'/selfList'共用一个处理函数
	group.POST("/project_template", h.projectTemplate)
	group.POST("/project/save", h.projectSave)
	group.POST("/project/read", h.readProject)
	group.POST("/project/recycle", h.recycleProject)
	group.POST("/project/recovery", h.recoveryProject)
	group.POST("/project_collect/collect", h.collectProject)
	group.POST("/project/edit", h.editProject)
	group.POST("/project/getLogBySelfProject", h.getLogBySelfProject)
	group.POST("/node", h.nodeList)

	t := NewTask()
	group.POST("/task_stages", t.taskStages)
	group.POST("/project_member/index", t.memberProjectList)
	group.POST("/task_stages/tasks", t.taskList)
	group.POST("/task/save", t.saveTask)
	group.POST("/task/sort", t.taskSort)
	group.POST("/task/selfList", t.myTaskList)
	group.POST("/task/read", t.readTask)
	group.POST("/task_member", t.listTaskMember)
	group.POST("/task/taskLog", t.taskLog)
	group.POST("/task/_taskWorkTimeList", t.taskWorkTimeList)
	group.POST("/task/saveTaskWorkTime", t.saveTaskWorkTime)
	group.POST("/file/uploadFiles", t.uploadFiles)
	group.POST("/task/taskSources", t.taskSources)
	group.POST("/task/createComment", t.createComment)

	a := NewAccount()
	group.POST("/account", a.account)

	d := NewDepartment()
	group.POST("/department", d.department)
	group.POST("/department/save", d.save)
	group.POST("/department/read", d.read)

	auth := NewAuth()
	group.POST("/auth", auth.authList)
	group.POST("/auth/apply", auth.apply)

	menu := NewMenu()
	group.POST("/menu/menu", menu.menuList)
}
