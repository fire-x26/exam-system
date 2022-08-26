package oj

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/model/common/response"
	ojReq "github.com/prl26/exam-system/server/model/oj/request"
	ojResp "github.com/prl26/exam-system/server/model/oj/response"
	"github.com/prl26/exam-system/server/service"
	"github.com/prl26/exam-system/server/utils"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/26 20:53

 * @Note:

 **/

type ProgrammApi struct{}

var cLanguageService = service.ServiceGroupApp.OjServiceServiceGroup.CLanguageService
var commonService = service.ServiceGroupApp.OjServiceServiceGroup.CommonService

// CheckProgramm 检验编程题
// @Tags OJ
// @Summary 检验编程题
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CheckProgramm true "检验编程题"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /oj/programm/check [post]
func (*ProgrammApi) CheckProgramm(c *gin.Context) {
	var r ojReq.CheckProgramm
	_ = c.ShouldBindJSON(&r)
	verify := utils.Rules{
		"Id":         {utils.NotEmpty()},
		"Code":       {utils.NotEmpty()},
		"LanguageId": {utils.NotEmpty()},
	}
	if err := utils.Verify(r, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	cases, err := commonService.FindProgrammCase(r.Id, r.LanguageId)
	if err != nil {
		response.FailWithMessage("未找到该题目或者该题目不支持该语言", c)
		return
	}
	switch r.LanguageId {
	case 1:
		result, err := cLanguageService.Check(r.Code, cases)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		response.OkWithData(result, c)
		return
	default:
		response.FailWithMessage("不支持该语言", c)
		return
	}
}

// Compile 编程题编译
// @Tags OJ
// @Summary	编程题编译
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CheckProgramm true "编程题编译"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /oj/programm/compile [post]
func (*ProgrammApi) Compile(c *gin.Context) {
	var r ojReq.Compile
	_ = c.ShouldBindJSON(&r)
	verify := utils.Rules{
		"Code":       {utils.NotEmpty()},
		"LanguageId": {utils.NotEmpty()},
	}
	if err := utils.Verify(r, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	switch r.LanguageId {
	case 1:
		fileId, expireTime, err := cLanguageService.Compile(r.Code)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		response.OkWithData(ojResp.Compile{FileId: fileId, ExpirationTime: *expireTime}, c)
		return
	default:
		response.FailWithMessage("不支持该语言", c)
		return
	}
}

// Execute 编程题执行
// @Tags OJ
// @Summary	编程题执行
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.Execute true "编程题执行"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /oj/programm/execute [post]
func (*ProgrammApi) Execute(c *gin.Context) {
	var r ojReq.Execute
	_ = c.ShouldBindJSON(&r)
	verify := utils.Rules{
		"FileId":     {utils.NotEmpty()},
		"LanguageId": {utils.NotEmpty()},
	}
	if err := utils.Verify(r, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	switch r.LanguageId {
	case 1:
		output, executeSituation, err := cLanguageService.Execute(r.FileId, r.Input, &r.ProgrammLimit)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		response.OkWithData(ojResp.Execute{Output: output, ExecuteSituation: *executeSituation}, c)
		return
	default:
		response.FailWithMessage("不支持该语言", c)
		return
	}
}
