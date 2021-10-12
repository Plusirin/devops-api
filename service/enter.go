package service

import (
	"devops-api/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
