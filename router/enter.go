package router

import (
	"devops-api/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
