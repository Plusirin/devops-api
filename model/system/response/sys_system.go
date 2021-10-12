package response

import "devops-api/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
