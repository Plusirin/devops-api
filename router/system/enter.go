package system

type RouterGroup struct {
	ApiRouter
	AuthorityRouter
	BaseRouter
	CasbinRouter
	InitRouter
	JwtRouter
	MenuRouter
	OperationRecordRouter
	SysRouter
	UserRouter
	ProjectRouter
	ToolRouter
}
