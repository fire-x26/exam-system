// 自动生成模板QuestionBankSupplyBlank
package questionBank

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// QuestionBankSupplyBlank 结构体
type SupplyBlank struct {
	global.GVA_MODEL
	Describe string `json:"describe" form:"describe" gorm:"column:describe;comment:填空题描述;"`
	IsOrder  *int   `json:"isOrder" form:"isOrder" gorm:"column:is_order;comment:是否要求有序;"`
	Num      *int   `json:"num" form:"num" gorm:"column:num;comment:可填项;"`
}

// TableName QuestionBankSupplyBlank 表名
func (SupplyBlank) TableName() string {
	return "les_questionBank_supply_blank"
}