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
//@function: CreateProject
//@description: 创建项目
//@param: u model.SysProject
//@return: err error, projectInter model.SysProject

type ProjectService struct {
}

func (projectService *ProjectService) CreateProject(p system.SysProject) (err error, projectInter system.SysProject) {
	var project system.SysProject
	if !errors.Is(global.GVA_DB.Where("project_name = ?", p.ProjectName).First(&project).Error, gorm.ErrRecordNotFound) { // 判断项目名是否存在
		return errors.New("项目名已存在"), projectInter
	}
	p.UUID = uuid.NewV4()
	err = global.GVA_DB.Create(&p).Error
	return err, p
}

//@author: [yunwei](yunwei@xwsoft.com.cn)
//@function: DeleteProject
//@description: 删除项目
//@param: id float64
//@return: err error

func (projectService *ProjectService) DeleteProject(id float64) (err error) {
	var project system.SysProject
	err = global.GVA_DB.Where("id = ?", id).Delete(&project).Error
	return
}

//@author: [yunwei](yunwei@xwsoft.com.cn)
//@function: ChangeProjectName
//@description: 修改项目名称
//@param: p *system.SysProject, newProjectName string
//@return: err error, projectInter *system.SysProject

func (projectService *ProjectService) ChangeProjectName(p *system.SysProject, newProjectName string) (err error, projectInter *system.SysProject) {
	var project system.SysProject
	err = global.GVA_DB.Where("id = ?", p.ID).First(&project).Update("project_name", newProjectName).Error
	return err, p
}

//@author: [yunwei](yunwei@xwsoft.com.cn)
//@function: ChangeProjectNickName
//@description: 修改项目别名
//@param: p *system.SysProject, newNickName string
//@return: err error, projectInter *system.SysProject

func (projectService *ProjectService) ChangeProjectNickName(p *system.SysProject, newNickName string) (err error, projectInter *system.SysProject) {
	var project system.SysProject
	err = global.GVA_DB.Where("id = ?", p.ID).First(&project).Update("nick_name", newNickName).Error
	return err, p
}

//@author: [yunwei](yunwei@xwsoft.com.cn)
//@function: ChangeProjectDescription
//@description: 修改项目描述
//@param: p *system.SysProject, newNickName string
//@return: err error, projectInter *system.SysProject

func (projectService *ProjectService) ChangeProjectDescription(p *system.SysProject, newDescription string) (err error, projectInter *system.SysProject) {
	var project system.SysProject
	err = global.GVA_DB.Where("id = ?", p.ID).First(&project).Update("description", newDescription).Error
	return err, p
}

//@author: [yunpei](yunwei@xwsoft.com.cn)
//@function: GetProjectInfoList
//@description: 分页获取项目数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func (projectService *ProjectService) GetProjectInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysProject{})
	var projectList []system.SysProject
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&projectList).Error
	return err, projectList, total
}

//@author: [yunwei](yunwei@xwsoft.com.cn)
//@function: SetProjectInfo
//@description: 设置项目信息
//@param: reqProject system.SysProject
//@return: err error, project system.SysProject

func (projectService *ProjectService) SetProjectInfo(reqProject system.SysProject) (err error, project system.SysProject) {
	err = global.GVA_DB.Updates(&reqProject).Error
	return err, reqProject
}

//@author: [yunwei](yunwei@xwsoft.com.cn)
//@function: GetProjectInfo
//@description: 获取项目信息
//@param: uuid uuid.UUID
//@return: err error, project system.SysProject

func (projectService *ProjectService) GetProjectInfo(id float64) (err error, project system.SysProject) {
	var reqProject system.SysProject
	err = global.GVA_DB.First(&reqProject, "id = ?", id).Error
	return err, reqProject
}

//@author: [yunwei](yunwei@xwsoft.com.cn)
//@function: FindProjectById
//@description: 通过id获取项目信息
//@param: id int
//@return: err error, project *model.SysProject

func (projectService *ProjectService) FindProjectById(id int) (err error, project *system.SysProject) {
	var p system.SysProject
	err = global.GVA_DB.Where("`id` = ?", id).First(&p).Error
	return err, &p
}
