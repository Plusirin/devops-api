package response

import (
	"devops-api/model/system"
)

type SysToolResponse struct {
	Tool system.SysTool `json:"tool"`
}
