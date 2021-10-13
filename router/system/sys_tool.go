package system

import (
	v1 "devops-api/api/v1"
	"devops-api/middleware"
	"github.com/gin-gonic/gin"
)

type ToolRouter struct {
}

func (s *ToolRouter) InitToolRouter(Router *gin.RouterGroup) {
	toolRouter := Router.Group("tool").Use(middleware.OperationRecord())
	//toolRouterWithoutRecord := Router.Group("tool")
	var toolApi = v1.ApiGroupApp.SystemApiGroup.ToolApi
	{
		toolRouter.POST("createTool", toolApi.CreateTool)
		toolRouter.POST("changeToolName", toolApi.ChangeToolName)
		toolRouter.POST("getToolInfo", toolApi.GetToolInfo)
		toolRouter.POST("getJobLastNum", toolApi.GetJobLastNum)
		toolRouter.DELETE("deleteTool", toolApi.DeleteTool)
	}
	{
		//toolRouterWithoutRecord.POST("getToolList", toolApi.GetToolList)
		//toolRouterWithoutRecord.GET("getToolInfo", toolApi.GetToolInfo)
	}
}
