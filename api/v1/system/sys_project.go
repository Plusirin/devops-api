package system

import (
	"devops-api/global"
	"devops-api/model/common/request"
	"devops-api/model/common/response"
	"devops-api/model/system"
	systemReq "devops-api/model/system/request"
	systemRes "devops-api/model/system/response"
	"devops-api/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProjectApi struct {
}

// @Tags SysProject
// @Summary 创建项目
// @accept application/json
// @Produce  application/json
// @Param data body systemReq.CreateProject true "项目名称, 别名, 描述"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"项目创建成功"}"
// @Router /project/createProject [post]
func (b *ProjectApi) CreateProject(c *gin.Context) {
	var p systemReq.CreateProject
	_ = c.ShouldBindJSON(&p)
	if err := utils.Verify(p, utils.CreateProjectVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	project := &system.SysProject{ProjectName: p.ProjectName,NickName: p.NickName,Description: p.Description}
	err, projectReturn := projectService.CreateProject(*project)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithDetailed(systemRes.SysProjectResponse{Project: projectReturn}, "创建失败", c)
	} else {
		response.OkWithDetailed(systemRes.SysProjectResponse{Project: projectReturn}, "创建成功", c)
	}
}

// @Tags SysProject
// @Summary 获取项目信息
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "项目ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /project/getProjectInfo [post]
func (b *ProjectApi) GetProjectInfo(c *gin.Context) {
	var reqId request.GetById
	_ = c.ShouldBindJSON(&reqId)
	if err := utils.Verify(reqId, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, ReqProject := projectService.GetProjectInfo(reqId.ID); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"userInfo": ReqProject}, "获取成功", c)
	}
}


// @Tags SysProject
// @Summary 删除项目
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "项目ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /project/deleteProject [delete]
func (b *ProjectApi) DeleteProject(c *gin.Context) {
	var reqId request.GetById
	_ = c.ShouldBindJSON(&reqId)
	if err := utils.Verify(reqId, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := projectService.DeleteProject(reqId.ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SysProject
// @Summary 修改项目名
// @accept application/json
// @Produce  application/json
// @Param data body systemReq.ChangeProjectName true "用户名, 原密码, 新密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /project/changeProjectName [post]
func (b *ProjectApi) ChangeProjectName(c *gin.Context) {
	var project systemReq.ChangeProjectName
	_ = c.ShouldBindJSON(&project)
	if err := utils.Verify(project, utils.ChangeProjectNameVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	p := &system.SysProject{GVA_MODEL: global.GVA_MODEL{ID: project.ID}, ProjectName: project.NewProjectName}
	if err, _ := projectService.ChangeProjectName(p, project.NewProjectName); err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage("修改失败", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}
//
//// @Tags SysUser
//// @Summary 分页获取用户列表
//// @Security ApiKeyAuth
//// @accept application/json
//// @Produce application/json
//// @Param data body request.PageInfo true "页码, 每页大小"
//// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
//// @Router /user/getUserList [post]
//func (b *ProjectApi) GetUserList(c *gin.Context) {
//	var pageInfo request.PageInfo
//	_ = c.ShouldBindJSON(&pageInfo)
//	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
//		response.FailWithMessage(err.Error(), c)
//		return
//	}
//	if err, list, total := userService.GetUserInfoList(pageInfo); err != nil {
//		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
//		response.FailWithMessage("获取失败", c)
//	} else {
//		response.OkWithDetailed(response.PageResult{
//			List:     list,
//			Total:    total,
//			Page:     pageInfo.Page,
//			PageSize: pageInfo.PageSize,
//		}, "获取成功", c)
//	}
//}
//
//// @Tags SysUser
//// @Summary 更改用户权限
//// @Security ApiKeyAuth
//// @accept application/json
//// @Produce application/json
//// @Param data body systemReq.SetUserAuth true "用户UUID, 角色ID"
//// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
//// @Router /user/setUserAuthority [post]
//func (b *ProjectApi) SetUserAuthority(c *gin.Context) {
//	var sua systemReq.SetUserAuth
//	_ = c.ShouldBindJSON(&sua)
//	if UserVerifyErr := utils.Verify(sua, utils.SetUserAuthorityVerify); UserVerifyErr != nil {
//		response.FailWithMessage(UserVerifyErr.Error(), c)
//		return
//	}
//	userID := utils.GetUserID(c)
//	uuid := utils.GetUserUuid(c)
//	if err := userService.SetUserAuthority(userID, uuid, sua.AuthorityId); err != nil {
//		global.GVA_LOG.Error("修改失败!", zap.Any("err", err))
//		response.FailWithMessage(err.Error(), c)
//	} else {
//		claims := utils.GetUserInfo(c)
//		j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
//		claims.AuthorityId = sua.AuthorityId
//		if token, err := j.CreateToken(*claims); err != nil {
//			global.GVA_LOG.Error("修改失败!", zap.Any("err", err))
//			response.FailWithMessage(err.Error(), c)
//		} else {
//			c.Header("new-token", token)
//			c.Header("new-expires-at", strconv.FormatInt(claims.ExpiresAt, 10))
//			response.OkWithMessage("修改成功", c)
//		}
//
//	}
//}
//
//// @Tags SysUser
//// @Summary 设置用户权限
//// @Security ApiKeyAuth
//// @accept application/json
//// @Produce application/json
//// @Param data body systemReq.SetUserAuthorities true "用户UUID, 角色ID"
//// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
//// @Router /user/setUserAuthorities [post]
//func (b *ProjectApi) SetUserAuthorities(c *gin.Context) {
//	var sua systemReq.SetUserAuthorities
//	_ = c.ShouldBindJSON(&sua)
//	if err := userService.SetUserAuthorities(sua.ID, sua.AuthorityIds); err != nil {
//		global.GVA_LOG.Error("修改失败!", zap.Any("err", err))
//		response.FailWithMessage("修改失败", c)
//	} else {
//		response.OkWithMessage("修改成功", c)
//	}
//}
//
//
//// @Tags SysUser
//// @Summary 设置用户信息
//// @Security ApiKeyAuth
//// @accept application/json
//// @Produce application/json
//// @Param data body system.SysUser true "ID, 用户名, 昵称, 头像链接"
//// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
//// @Router /user/setUserInfo [put]
//func (b *ProjectApi) SetUserInfo(c *gin.Context) {
//	var user system.SysUser
//	_ = c.ShouldBindJSON(&user)
//	if err := utils.Verify(user, utils.IdVerify); err != nil {
//		response.FailWithMessage(err.Error(), c)
//		return
//	}
//	if err, ReqUser := userService.SetUserInfo(user); err != nil {
//		global.GVA_LOG.Error("设置失败!", zap.Any("err", err))
//		response.FailWithMessage("设置失败", c)
//	} else {
//		response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "设置成功", c)
//	}
//}
//
