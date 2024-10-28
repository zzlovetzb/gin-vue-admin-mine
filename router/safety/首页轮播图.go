package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SwiperRouter struct {}

// InitSwiperRouter 初始化 轮播图 路由信息
func (s *SwiperRouter) InitSwiperRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	index_swiperRouter := Router.Group("index_swiper").Use(middleware.OperationRecord())
	index_swiperRouterWithoutRecord := Router.Group("index_swiper")
	index_swiperRouterWithoutAuth := PublicRouter.Group("index_swiper")
	{
		index_swiperRouter.POST("createSwiper", index_swiperApi.CreateSwiper)   // 新建轮播图
		index_swiperRouter.DELETE("deleteSwiper", index_swiperApi.DeleteSwiper) // 删除轮播图
		index_swiperRouter.DELETE("deleteSwiperByIds", index_swiperApi.DeleteSwiperByIds) // 批量删除轮播图
		index_swiperRouter.PUT("updateSwiper", index_swiperApi.UpdateSwiper)    // 更新轮播图
	}
	{
		index_swiperRouterWithoutRecord.GET("findSwiper", index_swiperApi.FindSwiper)        // 根据ID获取轮播图
		index_swiperRouterWithoutRecord.GET("getSwiperList", index_swiperApi.GetSwiperList)  // 获取轮播图列表
	}
	{
	    index_swiperRouterWithoutAuth.GET("getSwiperPublic", index_swiperApi.GetSwiperPublic)  // 轮播图开放接口
	}
}
