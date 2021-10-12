package system

import (
	"devops-api/global"
	"github.com/satori/go.uuid"
)

type SysTool struct {
	global.GVA_MODEL
	UUID        uuid.UUID `json:"uuid" gorm:"comment:项目UUID"`    // 工具UUID
	ToolName    string    `json:"toolName" gorm:"comment:工具名称"`  // 工具名称
	URL         string    `json:"url" gorm:"comment:url地址"`      // url地址
	Username    string    `json:"userName" gorm:"comment:用户名"`   // 用户名
	Token       string    `json:"token" gorm:"comment:token"`    // token
	Description string    `json:"description" gorm:"comment:描述"` // 描述
}
