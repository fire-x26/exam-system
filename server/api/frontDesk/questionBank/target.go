package questionBank

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/model/questionBank/enum/questionType"
	questionBankReq "github.com/prl26/exam-system/server/model/questionBank/vo/request"
	questionBankResp "github.com/prl26/exam-system/server/model/questionBank/vo/response"
	"github.com/prl26/exam-system/server/model/teachplan"
	"github.com/prl26/exam-system/server/service"
	"github.com/prl26/exam-system/server/utils"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type TargetApi struct {
}

var targetService = service.ServiceGroupApp.QuestionBankServiceGroup.TargetService
var targetOjService = service.ServiceGroupApp.QuestionBankServiceGroup.OjService.TargetService

func (*TargetApi) BeginPractice(c *gin.Context) {
	query := c.Query("lessonId")
	idInt, err := strconv.Atoi(query)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	lessonId := uint(idInt)
	detail, err := lessonService.FindLessonDetail(lessonId, true)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	studentId := utils.GetStudentId(c)
	go func() {
		practiceService.UpdatePracticeRecord(lessonId, studentId)
		now := time.Now()
		ip := c.ClientIP()
		r := &teachplan.PracticeRecord{
			LessonId:  lessonId,
			StudentId: studentId,
			BeginTime: now,
			BeginIp:   ip,
		}
		practiceService.CreatePracticeRecord(r)
	}()
	response.OkWithData(detail, c)
}

func (*TargetApi) FindTargetByKnowledgeId(c *gin.Context) {
	search := questionBankReq.TargetPracticeSearch{}
	_ = c.ShouldBindQuery(&search)
	verify := utils.Rules{
		"Page":     {utils.NotEmpty()},
		"PageSize": {utils.NotEmpty()},
	}
	if err := utils.Verify(search, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := targetService.FindTargetPracticeList(search.TargetPracticeCriteria, search.PageInfo)
	uints := make([]uint, len(list))
	for i, practice := range list {
		uints[i] = practice.ID
	}
	studentId := utils.GetStudentId(c)
	answer := practiceService.FindHistoryAnswer(questionType.Target, uints, studentId)
	for _, practice := range list {
		practice.IsDone = answer.History[practice.ID].Exist
		practice.HistoryScore = answer.History[practice.ID].Score
	}
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     search.Page,
			PageSize: search.PageSize,
		}, "获取成功", c)
	}
}

func (*TargetApi) FindTargetDetail(c *gin.Context) {
	query := c.Query("id")
	idInt, err := strconv.Atoi(query)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	targetId := uint(idInt)
	detail, err := targetService.FindDetail(targetId, false)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, fmt.Errorf("获取失败:%s", err.Error()))
		return
	} else {
		studentId := utils.GetStudentId(c)
		address, isGenerateAddress := targetService.QueryPracticeRecord(studentId, targetId)
		q := &questionBankResp.TargetDetail{
			TargetDetail:      detail,
			IsGenerateAddress: isGenerateAddress,
			Address:           address,
		}
		history, isDone := targetService.QueryHistory(studentId, targetId)
		q.HistoryScore = history
		if isDone {
			q.IsDone = true
		}
		questionBankResp.OkWithDetailed(q, "获取成功", c)
		return
	}
}

func (*TargetApi) PracticeGenerateInstance(c *gin.Context) {
	query := c.Query("id")
	idInt, err := strconv.Atoi(query)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	targetId := uint(idInt)
	byteCodeModel := targetService.GetByteCode(targetId)
	if byteCodeModel == nil {
		questionBankResp.NotFind(c)
		return
	}
	salt, address, deployAddress, err := targetOjService.GenerateInstance(byteCodeModel.ByteCode)
	if err != nil {
		questionBankResp.ErrorHandle(c, fmt.Errorf("该题生成实例错误，请联系管理员检测"))
		return
	}
	studentId := utils.GetStudentId(c)
	targetService.PracticeRecord(studentId, targetId, address)
	questionBankResp.OkWithDetailed(questionBankResp.TargetGenerateInstance{
		Address: address,
		//ByteCode: byteCodeModel.ByteCode,
		DeployAddress: deployAddress,
		Salt:          salt,
	}, "生成成功", c)
}

func (*TargetApi) PracticeScore(c *gin.Context) {
	query := c.Query("id")
	idInt, err := strconv.Atoi(query)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	targetId := uint(idInt)
	studentId := utils.GetStudentId(c)
	address, isGenerateAddress := targetService.QueryPracticeRecord(studentId, targetId)
	if !isGenerateAddress {
		questionBankResp.CheckHandle(c, fmt.Errorf("暂未生成实例地址"))
		return
	}
	score, err := targetOjService.QueryScore(address)
	if err != nil {
		questionBankResp.ErrorHandle(c, fmt.Errorf("获取分数错误，请联系管理员或重新生成实例"))
		return
	}
	go func() {
		//t := practiceService.FindTheLatestRecord(lessonId, studentId)
		practiceService.CreatePracticeItem(questionType.Target, targetId, 25, studentId, uint(score))
		practiceService.UpdatePracticeAnswer(questionType.Target, targetId, studentId, uint(score))
	}()
	questionBankResp.OkWithDetailed(score, "获取成功", c)
}
