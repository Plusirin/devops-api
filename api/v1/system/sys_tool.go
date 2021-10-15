package system

import (
	"devops-api/global"
	"devops-api/model/common/request"
	"devops-api/model/common/response"
	"devops-api/model/system"
	systemReq "devops-api/model/system/request"
	systemRes "devops-api/model/system/response"
	"devops-api/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type ToolApi struct {
}

// @Tags SysTool
// @Summary 创建工具
// @accept application/json
// @Produce  application/json
// @Param data body systemReq.CreateTool true "工具名称, 别名, 描述"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"工具创建成功"}"
// @Router /tool/createTool [post]
func (b *ToolApi) CreateTool(c *gin.Context) {
	var p systemReq.CreateTool
	_ = c.ShouldBindJSON(&p)
	if err := utils.Verify(p, utils.CreateToolVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	tool := &system.SysTool{
		ToolName:    p.ToolName,
		URL:         p.URL,
		Username:    p.Username,
		Token:       p.Token,
		Description: p.Description,
	}
	err, toolReturn := toolService.CreateTool(*tool)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithDetailed(systemRes.SysToolResponse{Tool: toolReturn}, "创建失败", c)
	} else {
		response.OkWithDetailed(systemRes.SysToolResponse{Tool: toolReturn}, "创建成功", c)
	}
}

// @Tags SysTool
// @Summary 获取工具信息
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "工具ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tool/getToolInfo [post]
func (b *ToolApi) GetToolInfo(c *gin.Context) {
	var reqId request.GetById
	_ = c.ShouldBindJSON(&reqId)
	if err := utils.Verify(reqId, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, ReqTool := toolService.GetToolInfo(reqId.ID); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"toolInfo": ReqTool}, "获取成功", c)
	}
}

// @Tags SysTool
// @Summary 删除工具
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "工具ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tool/deleteTool [delete]
func (b *ToolApi) DeleteTool(c *gin.Context) {
	var reqId request.GetById
	_ = c.ShouldBindJSON(&reqId)
	if err := utils.Verify(reqId, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := toolService.DeleteTool(reqId.ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SysTool
// @Summary 修改工具名
// @accept application/json
// @Produce  application/json
// @Param data body systemReq.ChangeToolName true "用户名, 原密码, 新密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /tool/changeToolName [post]
func (b *ToolApi) ChangeToolName(c *gin.Context) {
	var tool systemReq.ChangeToolName
	_ = c.ShouldBindJSON(&tool)
	if err := utils.Verify(tool, utils.ChangeToolNameVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	p := &system.SysTool{GVA_MODEL: global.GVA_MODEL{ID: tool.ID}, ToolName: tool.NewToolName}
	if err, _ := toolService.ChangeToolName(p, tool.NewToolName); err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage("修改失败", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// @Tags SysTool
// @Summary 获取任务最后编号
// @accept application/json
// @Produce  application/json
// @Param data body systemReq.GetJobLastNum true "id, jobName"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tool/getJobLastNum [post]
func (b *ToolApi) GetJobLastNum(c *gin.Context) {
	var job systemReq.GetJobLastNum
	_ = c.ShouldBindJSON(&job)
	if err := utils.Verify(job, utils.GetJobLastNumVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err, ReqTool := toolService.GetJobLastNum(job.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	}

	// 定义指定任务url（使用结构体中参数进行拼接）
	JenkinsJobUrl := ReqTool.URL + "/job/" + job.JobName + "/lastBuild/buildNumber"
	client := &http.Client{Timeout: 5 * time.Second}
	// 使用newrequest准备http调用命令（）
	loginPara := fmt.Sprintf("username=%s&password=%s", ReqTool.Username, ReqTool.Token)
	req, errRq := http.NewRequest("GET", JenkinsJobUrl, strings.NewReader(loginPara))
	if errRq != nil {
		global.GVA_LOG.Error("Jenkins Show Last Number Fail!", zap.Any("err", err))
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, errDo := client.Do(req)
	if errDo == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			global.GVA_LOG.Error("Jenkins Show Last Number Fail")
			response.FailWithMessage("获取失败", c)
		}
		global.GVA_LOG.Info("Jenkins Show Last Number")
		// 返回查询结果正文，正文内容为执行编号（字符串格式）
		global.GVA_LOG.Info("Jenkins Show Last Number  Finish")
		response.OkWithDetailed(gin.H{"toolInfo": string(body)}, "获取成功", c)
	} else {
		global.GVA_LOG.Error("Jenkins Show Last Number Fail", zap.Any("err", err))
	}
}

// @Tags SysTool
// @Summary 获取任务最后状态
// @accept application/json
// @Produce  application/json
// @Param data body systemReq.GetJobLastNum true "id, jobName"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tool/getJobLastState [post]
func (b *ToolApi) GetJobLastState(c *gin.Context) {
	var job systemReq.GetJobLastNum
	_ = c.ShouldBindJSON(&job)
	if err := utils.Verify(job, utils.GetJobLastNumVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err, ReqTool := toolService.GetJobLastNum(job.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	}

	// 定义指定任务url（使用结构体中参数进行拼接）
	JenkinsJobUrl := ReqTool.URL + "/job/" + job.JobName + "/lastBuild/api/json"
	client := &http.Client{Timeout: 5 * time.Second}
	// 使用newrequest准备http调用命令（）
	loginPara := fmt.Sprintf("username=%s&password=%s", ReqTool.Username, ReqTool.Token)
	req, errRq := http.NewRequest("GET", JenkinsJobUrl, strings.NewReader(loginPara))
	if errRq != nil {
		global.GVA_LOG.Error("Jenkins Show Last State Fail!", zap.Any("err", err))
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, errDo := client.Do(req)
	if errDo == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			global.GVA_LOG.Error("Jenkins Show Last State Fail")
			response.FailWithMessage("获取失败", c)
		}
		global.GVA_LOG.Info("Jenkins Show Last State")
		// 返回查询结果正文，正文内容为执行编号（字符串格式）
		global.GVA_LOG.Info("Jenkins Show Last State  Finish")
		response.OkWithDetailed(gin.H{"toolInfo": string(body)}, "获取成功", c)
	} else {
		global.GVA_LOG.Error("Jenkins Show Last State Fail", zap.Any("err", err))
	}
}

// @Tags SysTool
// @Summary 更具buildnum获取console
// @accept application/json
// @Produce  application/json
// @Param data body systemReq.GetJobLastNum true "id, jobName"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tool/getJobConsoleWithNum [post]
func (b *ToolApi) GetJobConsoleWithNum(c *gin.Context) {
	var job systemReq.GetJobConsole
	_ = c.ShouldBindJSON(&job)
	if err := utils.Verify(job, utils.GetJobLastNumVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err, ReqTool := toolService.GetJobLastNum(job.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	}

	// 定义指定任务url（使用结构体中参数进行拼接）
	JenkinsJobUrl := ReqTool.URL + "/job/" + job.JobName + "/" + job.JobID + "/logText/progressiveText"
	client := &http.Client{Timeout: 5 * time.Second}
	// 使用newrequest准备http调用命令（）
	loginPara := fmt.Sprintf("username=%s&password=%s", ReqTool.Username, ReqTool.Token)
	req, errRq := http.NewRequest("GET", JenkinsJobUrl, strings.NewReader(loginPara))
	if errRq != nil {
		global.GVA_LOG.Error("Jenkins Show Console Fail!", zap.Any("err", err))
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, errDo := client.Do(req)
	if errDo == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			global.GVA_LOG.Error("Jenkins Show Console Fail")
			response.FailWithMessage("获取失败", c)
		}
		global.GVA_LOG.Info("Jenkins Show Console Finish")
		response.OkWithDetailed(gin.H{"toolInfo": string(body)}, "获取成功", c)
	} else {
		global.GVA_LOG.Error("Jenkins Show Console Fail", zap.Any("err", err))
	}
}
