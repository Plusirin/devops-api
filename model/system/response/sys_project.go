package response


import (
	"devops-api/model/system"
)

type SysProjectResponse struct {
	Project system.SysProject `json:"project"`
}