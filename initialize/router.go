package initialize

import (
	_ "devops-api/docs"
	"devops-api/global"
	"devops-api/middleware"
	"devops-api/router"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()

	global.GVA_LOG.Info("use middleware logger")
	// 跨域
	Router.Use(middleware.Cors()) // 如需跨域可以打开
	global.GVA_LOG.Info("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")

	//获取路由组实例
	systemRouter := router.RouterGroupApp.System
	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		systemRouter.InitInitRouter(PublicGroup) // 自动初始化相关
	}
	PrivateGroup := Router.Group("")
	//PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		systemRouter.InitApiRouter(PrivateGroup)                    // 注册功能api路由
		systemRouter.InitJwtRouter(PrivateGroup)                    // jwt相关路由
		systemRouter.InitUserRouter(PrivateGroup)                   // 注册用户路由
		systemRouter.InitMenuRouter(PrivateGroup)                   // 注册menu路由
		systemRouter.InitSystemRouter(PrivateGroup)                 // system相关路由
		systemRouter.InitCasbinRouter(PrivateGroup)                 // 权限相关路由
		systemRouter.InitAuthorityRouter(PrivateGroup)              // 注册角色路由
		systemRouter.InitSysOperationRecordRouter(PrivateGroup)     // 操作记录
		systemRouter.InitProjectRouter(PrivateGroup)     // 项目路由
	}
	
	global.GVA_LOG.Info("router register success")
	return Router
}
