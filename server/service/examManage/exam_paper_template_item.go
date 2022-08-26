package examManage

import (
	"fmt"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/examManage"
	examManageReq "github.com/prl26/exam-system/server/model/examManage/request"
	"github.com/prl26/exam-system/server/model/questionBank"
	"math/rand"
)

type PaperTemplateItemService struct {
}

// CreatePaperTemplateItem 创建PaperTemplateItem记录
// Author [piexlmax](https://github.com/piexlmax)
func (paperTemplateItemService *PaperTemplateItemService) CreatePaperTemplateItem(paperTemplateItem examManage.PaperTemplateItem) (err error) {
	err = global.GVA_DB.Create(&paperTemplateItem).Error
	return err
}

// DeletePaperTemplateItem 删除PaperTemplateItem记录
// Author [piexlmax](https://github.com/piexlmax)
func (paperTemplateItemService *PaperTemplateItemService) DeletePaperTemplateItem(paperTemplateItem examManage.PaperTemplateItem) (err error) {
	err = global.GVA_DB.Delete(&paperTemplateItem).Error
	return err
}

// DeletePaperTemplateItemByIds 批量删除PaperTemplateItem记录
// Author [piexlmax](https://github.com/piexlmax)
func (paperTemplateItemService *PaperTemplateItemService) DeletePaperTemplateItemByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]examManage.PaperTemplateItem{}, "id in ?", ids.Ids).Error
	return err
}

// UpdatePaperTemplateItem 更新PaperTemplateItem记录
// Author [piexlmax](https://github.com/piexlmax)
func (paperTemplateItemService *PaperTemplateItemService) UpdatePaperTemplateItem(paperTemplateItem examManage.PaperTemplateItem) (err error) {
	err = global.GVA_DB.Save(&paperTemplateItem).Error
	return err
}

// GetPaperTemplateItem 根据id获取PaperTemplateItem记录
// Author [piexlmax](https://github.com/piexlmax)
func (paperTemplateItemService *PaperTemplateItemService) GetPaperTemplateItem(id uint) (paperTemplateItem examManage.PaperTemplateItem, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&paperTemplateItem).Error
	return
}

// GetPaperTemplateItemInfoList 分页获取PaperTemplateItem记录
// Author [piexlmax](https://github.com/piexlmax)
func (paperTemplateItemService *PaperTemplateItemService) GetPaperTemplateItemInfoList(info examManageReq.PaperTemplateItemSearch) (list []examManage.PaperTemplateItem, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&examManage.PaperTemplateItem{})
	var paperTemplateItems []examManage.PaperTemplateItem
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ChapterId != nil {
		db = db.Where("chapter = ?", info.ChapterId)
	}
	if info.ProblemType != nil {
		db = db.Where("problem_type = ?", info.ProblemType)
	}
	if info.QuestionType != nil {
		db = db.Where("difficulty = ?", info.QuestionType)
	}
	if info.Num != nil {
		db = db.Where("num = ?", info.Num)
	}
	if info.Score != nil {
		db = db.Where("score = ?", info.Score)
	}
	if info.TemplateId != nil {
		db = db.Where("template_id = ?", info.TemplateId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&paperTemplateItems).Error
	return paperTemplateItems, total, err
}
func (paperTemplateItemService *PaperTemplateItemService) GetPaperId(info examManage.PaperTemplateItem) (paperId uint) {
	var paper examManage.ExamPaper
	global.GVA_DB.Where("id = ?", info.TemplateId).Find(&paper)
	return paper.ID
}

func (paperTemplateItemService *PaperTemplateItemService) SetPaperChoiceQuestion(info []examManage.PaperTemplateItem) (err error) {
	var list []questionBank.MultipleChoice
	//fmt.Println(len())
	paperId := paperTemplateItemService.GetPaperId(info[0])
	for i := 0; i < len(info); i++ {
		num := info[i].Num
		uuid := rand.Intn(9)
		err = global.GVA_DB.Where("problem_type = ? and can_exam = ?", info[i].ProblemType, 1).Order(uuid).Limit(*num).Find(&list).Error
		if err != nil {
			return
		} else {
			if len(list) > 0 {
				for j := 0; j < *num; j++ {
					questionMerge := examManage.PaperQuestionMerge{
						GVA_MODEL:    global.GVA_MODEL{},
						PaperId:      &paperId,
						QuestionId:   &list[j].ID,
						Score:        info[i].Score,
						QuestionType: info[i].QuestionType,
						ProblemType:  info[i].ProblemType,
					}
					err = global.GVA_DB.Create(&questionMerge).Error
					if err != nil {
						return
					}
				}
			}
		}
	}
	return
}
func (paperTemplateItemService *PaperTemplateItemService) SetPaperJudgeQuestion(info []examManage.PaperTemplateItem) (err error) {
	var list []questionBank.Judge
	//fmt.Println(len())
	paperId := paperTemplateItemService.GetPaperId(info[0])
	for i := 0; i < len(info); i++ {
		num := info[i].Num
		uuid := rand.Intn(9)
		err = global.GVA_DB.Where("problem_type = ? and can_exam = ?", info[i].ProblemType, 1).Order(uuid).Limit(*num).Find(&list).Error
		if err != nil {
			return
		} else {
			if len(list) > 0 {
				for j := 0; j < *num; j++ {
					questionMerge := examManage.PaperQuestionMerge{
						GVA_MODEL:    global.GVA_MODEL{},
						PaperId:      &paperId,
						QuestionId:   &list[j].ID,
						Score:        info[i].Score,
						QuestionType: info[i].QuestionType,
						ProblemType:  info[i].ProblemType,
					}
					err = global.GVA_DB.Create(&questionMerge).Error
					if err != nil {
						return
					}
				}
			}
		}
	}
	return
}
func (paperTemplateItemService *PaperTemplateItemService) SetPaperBlankQuestion(info []examManage.PaperTemplateItem) (err error) {
	var list []questionBank.SupplyBlank
	paperId := paperTemplateItemService.GetPaperId(info[0])
	for i := 0; i < len(info); i++ {
		num := info[i].Num
		uuid := rand.Intn(9)
		err = global.GVA_DB.Where("problem_type = ? and can_exam = ?", info[i].ProblemType, 1).Order(uuid).Limit(*num).Find(&list).Error
		if err != nil {
			return
		} else {
			if len(list) > 0 {
				for j := 0; j < *num; j++ {
					fmt.Println(j)
					questionMerge := examManage.PaperQuestionMerge{
						GVA_MODEL:    global.GVA_MODEL{},
						PaperId:      &paperId,
						QuestionId:   &list[j].ID,
						Score:        info[i].Score,
						QuestionType: info[i].QuestionType,
						ProblemType:  info[i].ProblemType,
					}
					err = global.GVA_DB.Create(&questionMerge).Error
					if err != nil {
						return
					}
				}
			}
		}
	}
	return
}
func (paperTemplateItemService *PaperTemplateItemService) SetPaperProgramQuestion(info []examManage.PaperTemplateItem) (err error) {
	var list []questionBank.Programm
	//fmt.Println(len())
	paperId := paperTemplateItemService.GetPaperId(info[0])
	for i := 0; i < len(info); i++ {
		num := info[i].Num
		uuid := rand.Intn(9)
		err = global.GVA_DB.Where("problem_type = ? and can_exam = ?", info[i].ProblemType, 1).Order(uuid).Limit(*num).Find(&list).Error
		if err != nil {
			return
		} else {
			if len(list) > 0 {
				for j := 0; j < *num; j++ {
					questionMerge := examManage.PaperQuestionMerge{
						GVA_MODEL:    global.GVA_MODEL{},
						PaperId:      &paperId,
						QuestionId:   &list[j].ID,
						Score:        info[i].Score,
						QuestionType: info[i].QuestionType,
						ProblemType:  info[i].ProblemType,
					}
					err = global.GVA_DB.Create(&questionMerge).Error
					if err != nil {
						return
					}
				}
			}
		}
	}
	return
}

func (paperTemplateItemService *PaperTemplateItemService) SetPaperQuestion(info []examManage.PaperTemplateItem) (err error) {
	err = paperTemplateItemService.SetPaperChoiceQuestion(info)
	if err != nil {
		return
	} else {
		err = paperTemplateItemService.SetPaperJudgeQuestion(info)
		if err != nil {
			return
		} else {
			err = paperTemplateItemService.SetPaperBlankQuestion(info)
			if err != nil {
				return
			} else {
				err = paperTemplateItemService.SetPaperProgramQuestion(info)
				if err != nil {
					return
				}
			}
		}
	}
	return
}
