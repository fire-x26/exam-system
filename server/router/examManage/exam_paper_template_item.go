package examManage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PaperTemplateItemRouter struct {
}

// InitPaperTemplateItemRouter 初始化 PaperTemplateItem 路由信息
func (s *PaperTemplateItemRouter) InitPaperTemplateItemRouter(Router *gin.RouterGroup) {
	paperTemplateItemRouter := Router.Group("paperTemplateItem").Use(middleware.OperationRecord())
	paperTemplateItemRouterWithoutRecord := Router.Group("paperTemplateItem")
	var paperTemplateItemApi = v1.ApiGroupApp.ExammanageApiGroup.PaperTemplateItemApi
	{
		paperTemplateItemRouter.POST("createPaperTemplateItem", paperTemplateItemApi.CreatePaperTemplateItem)   // 新建PaperTemplateItem
		paperTemplateItemRouter.DELETE("deletePaperTemplateItem", paperTemplateItemApi.DeletePaperTemplateItem) // 删除PaperTemplateItem
		paperTemplateItemRouter.DELETE("deletePaperTemplateItemByIds", paperTemplateItemApi.DeletePaperTemplateItemByIds) // 批量删除PaperTemplateItem
		paperTemplateItemRouter.PUT("updatePaperTemplateItem", paperTemplateItemApi.UpdatePaperTemplateItem)    // 更新PaperTemplateItem
	}
	{
		paperTemplateItemRouterWithoutRecord.GET("findPaperTemplateItem", paperTemplateItemApi.FindPaperTemplateItem)        // 根据ID获取PaperTemplateItem
		paperTemplateItemRouterWithoutRecord.GET("getPaperTemplateItemList", paperTemplateItemApi.GetPaperTemplateItemList)  // 获取PaperTemplateItem列表
	}
}
