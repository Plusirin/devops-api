package system

type ServiceGroup struct {
	JwtService
	ApiService
	AuthorityService
	BaseMenuService
	CasbinService
	InitDBService
	MenuService
	OperationRecordService
	SystemConfigService
	UserService
	ProjectService
	ToolService
}
