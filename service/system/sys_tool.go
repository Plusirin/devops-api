package system

import (
	"errors"

	"devops-api/global"
	"devops-api/model/common/request"
	"devops-api/model/system"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

//@author: [yunwei](yunwei@xwsoft.com.cn)
//@function: CreateTool
//@description: 创建工具
//@param: u model.SysTool
//@return: err error, toolInter model.SysTool

type ToolService struct {
}

func (toolService *ToolService) CreateTool(p system.SysTool) (err error, toolInter system.SysTool) {
	var tool system.SysTool
	if !errors.Is(global.GVA_DB.Where("tool_name = ?", p.ToolName).First(&tool).Error, gorm.ErrRecordNotFound) { // 判断工具名是否存在
		return errors.New("工具名已存在"), toolInter
	}
	p.UUID = uuid.NewV4()
	err = global.GVA_DB.Create(&p).Error
	return err, p
}

//@author: [yunwei](yunwei@xwsoft.com.cn)
//@function: DeleteTool
//@description: 删除工具
//@param: id float64
//@return: err error

func (toolService *ToolService) DeleteTool(id float64) (err error) {
	var tool system.SysTool
	err = global.GVA_DB.Where("id = ?", id).Delete(&tool).Error
	return
}

//@author: [yunwei](yunwei@xwsoft.com.cn)
//@function: ChangeToolName
//@description: 修改工具名称
//@param: p *system.SysTool, newToolName string
//@return: err error, toolInter *system.SysTool

func (toolService *ToolService) ChangeToolName(p *system.SysTool, newToolName string) (err error, toolInter *system.SysTool) {
	var tool system.SysTool
	err = global.GVA_DB.Where("id = ?", p.ID).First(&tool).Update("tool_name", newToolName).Error
	return err, p
}

//@author: [yunwei](yunwei@xwsoft.com.cn)
//@function: ChangeToolNickName
//@description: 修改工具别名
//@param: p *system.SysTool, newNickName string
//@return: err error, toolInter *system.SysTool

func (toolService *ToolService) ChangeToolNickName(p *system.SysTool, newNickName string) (err error, toolInter *system.SysTool) {
	var tool system.SysTool
	err = global.GVA_DB.Where("id = ?", p.ID).First(&tool).Update("nick_name", newNickName).Error
	return err, p
}

//@author: [yunwei](yunwei@xwsoft.com.cn)
//@function: ChangeToolDescription
//@description: 修改工具描述
//@param: p *system.SysTool, newNickName string
//@return: err error, toolInter *system.SysTool

func (toolService *ToolService) ChangeToolDescription(p *system.SysTool, newDescription string) (err error, toolInter *system.SysTool) {
	var tool system.SysTool
	err = global.GVA_DB.Where("id = ?", p.ID).First(&tool).Update("description", newDescription).Error
	return err, p
}

//@author: [yunpei](yunwei@xwsoft.com.cn)
//@function: GetToolInfoList
//@description: 分页获取工具数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func (toolService *ToolService) GetToolInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysTool{})
	var toolList []system.SysTool
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&toolList).Error
	return err, toolList, total
}

//@author: [yunwei](yunwei@xwsoft.com.cn)
//@function: SetToolInfo
//@description: 设置工具信息
//@param: reqTool system.SysTool
//@return: err error, tool system.SysTool

func (toolService *ToolService) SetToolInfo(reqTool system.SysTool) (err error, tool system.SysTool) {
	err = global.GVA_DB.Updates(&reqTool).Error
	return err, reqTool
}

//@author: [yunwei](yunwei@xwsoft.com.cn)
//@function: GetToolInfo
//@description: 获取工具信息
//@param: id float64
//@return: err error, tool system.SysTool

func (toolService *ToolService) GetToolInfo(id float64) (err error, tool system.SysTool) {
	var reqTool system.SysTool
	err = global.GVA_DB.First(&reqTool, "id = ?", id).Error
	return err, reqTool
}

//@author: [yunwei](yunwei@xwsoft.com.cn)
//@function: FindToolById
//@description: 通过id获取工具信息
//@param: id int
//@return: err error, tool *model.SysTool

func (toolService *ToolService) FindToolById(id int) (err error, tool *system.SysTool) {
	var p system.SysTool
	err = global.GVA_DB.Where("`id` = ?", id).First(&p).Error
	return err, &p
}

//@author: [yunwei](yunwei@xwsoft.com.cn)
//@function: GetJobLastNum
//@description: 获取job最新的NUMBER
//@param: id float64
//@return: err error, tool system.SysTool

func (toolService *ToolService) GetJobLastNum(id uint) (err error, tool system.SysTool) {
	var reqTool system.SysTool
	err = global.GVA_DB.First(&reqTool, "id = ?", id).Error
	return err, reqTool
}
