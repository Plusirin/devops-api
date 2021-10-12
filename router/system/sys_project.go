package system

import (
	v1 "devops-api/api/v1"
	"devops-api/middleware"
	"github.com/gin-gonic/gin"
)

type ProjectRouter struct {
}

func (s *ProjectRouter) InitProjectRouter(Router *gin.RouterGroup) {
	projectRouter := Router.Group("project").Use(middleware.OperationRecord())
	//projectRouterWithoutRecord := Router.Group("project")
	var projectApi = v1.ApiGroupApp.SystemApiGroup.ProjectApi
	{
		projectRouter.POST("createProject", projectApi.CreateProject)                     // 用户注册账号
		projectRouter.POST("changeProjectName", projectApi.ChangeProjectName)         // 用户修改密码
		projectRouter.POST("getProjectInfo", projectApi.GetProjectInfo)         // 用户修改密码
		projectRouter.DELETE("deleteProject", projectApi.DeleteProject)               // 删除用户
	}
	{
		//projectRouterWithoutRecord.POST("getUserList", projectApi.GetUserList) // 分页获取用户列表
		//projectRouterWithoutRecord.GET("getUserInfo", projectApi.GetUserInfo)  // 获取自身信息
	}
}
