package system

import "devops-api/service"

type ApiGroup struct {
	SystemApiApi
	AuthorityApi
	BaseApi
	CasbinApi
	SystemApi
	DBApi
	JwtApi
	OperationRecordApi
	AuthorityMenuApi
	ProjectApi
	ToolApi
}

var authorityService = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
var apiService = service.ServiceGroupApp.SystemServiceGroup.ApiService
var menuService = service.ServiceGroupApp.SystemServiceGroup.MenuService
var casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService
var initDBService = service.ServiceGroupApp.SystemServiceGroup.InitDBService
var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService
var baseMenuService = service.ServiceGroupApp.SystemServiceGroup.BaseMenuService
var operationRecordService = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService
var userService = service.ServiceGroupApp.SystemServiceGroup.UserService
var systemConfigService = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService
var projectService = service.ServiceGroupApp.SystemServiceGroup.ProjectService
var toolService = service.ServiceGroupApp.SystemServiceGroup.ToolService
