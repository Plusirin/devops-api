package system

import (
	"devops-api/global"
	"github.com/satori/go.uuid"
)

type SysEnv struct {
	global.GVA_MODEL
	UUID        uuid.UUID `json:"uuid" gorm:"comment:项目UUID"`                    // 环境UUID
	EnvName     string    `json:"envName" gorm:"comment:环境名称"`                   // 环境名称
	NickName    string    `json:"nickName" gorm:"comment:别名"`                    // 别名
	URL         string    `json:"url" gorm:"comment:url地址"`                      // url地址
	AuthorityId string    `json:"authorityId" gorm:"default:888;comment:用户角色ID"` // 用户角色ID
	Description string    `json:"description" gorm:"comment:描述"`                 // 描述
}
