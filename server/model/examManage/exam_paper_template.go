// 自动生成模板ExamPaperTemplate
package examManage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ExamPaperTemplate 结构体
type ExamPaperTemplate struct {
	global.GVA_MODEL
	Course_id *int   `json:"course_id" form:"course_id" gorm:"column:course_id;comment:课程id;"`
	User_id   *int   `json:"user_id" form:"user_id" gorm:"column:user_id;comment:创建者id;"`
	Name      string `json:"name" form:"name" gorm:"column:name;comment:模板名称;"`
	Memo      string `json:"memo" form:"memo" gorm:"column:memo;comment:模板备注;"`
}

// TableName ExamPaperTemplate 表名
func (ExamPaperTemplate) TableName() string {
	return "exam_paper_template"
}