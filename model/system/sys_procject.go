package system

import (
	"devops-api/global"
	"github.com/satori/go.uuid"
)

type SysProject struct {
	global.GVA_MODEL
	UUID        uuid.UUID `json:"uuid" gorm:"comment:项目UUID"`              // 项目UUID
	ProjectName string    `json:"projectName" gorm:"comment:项目名称"`         // 项目名称
	NickName    string    `json:"nickName" gorm:"comment:别名"` // 别名
	Description string    `json:"description" gorm:"comment:描述"`           // 描述
}
