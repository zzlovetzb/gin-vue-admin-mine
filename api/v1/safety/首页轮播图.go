package safety

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/safety"
    safetyReq "github.com/flipped-aurora/gin-vue-admin/server/model/safety/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type SwiperApi struct {}



// CreateSwiper 创建轮播图
// @Tags Swiper
// @Summary 创建轮播图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.Swiper true "创建轮播图"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /index_swiper/createSwiper [post]
func (index_swiperApi *SwiperApi) CreateSwiper(c *gin.Context) {
	var index_swiper safety.Swiper
	err := c.ShouldBindJSON(&index_swiper)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = index_swiperService.CreateSwiper(&index_swiper)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteSwiper 删除轮播图
// @Tags Swiper
// @Summary 删除轮播图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.Swiper true "删除轮播图"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /index_swiper/deleteSwiper [delete]
func (index_swiperApi *SwiperApi) DeleteSwiper(c *gin.Context) {
	id := c.Query("id")
	err := index_swiperService.DeleteSwiper(id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteSwiperByIds 批量删除轮播图
// @Tags Swiper
// @Summary 批量删除轮播图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /index_swiper/deleteSwiperByIds [delete]
func (index_swiperApi *SwiperApi) DeleteSwiperByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	err := index_swiperService.DeleteSwiperByIds(ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateSwiper 更新轮播图
// @Tags Swiper
// @Summary 更新轮播图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.Swiper true "更新轮播图"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /index_swiper/updateSwiper [put]
func (index_swiperApi *SwiperApi) UpdateSwiper(c *gin.Context) {
	var index_swiper safety.Swiper
	err := c.ShouldBindJSON(&index_swiper)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = index_swiperService.UpdateSwiper(index_swiper)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSwiper 用id查询轮播图
// @Tags Swiper
// @Summary 用id查询轮播图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query safety.Swiper true "用id查询轮播图"
// @Success 200 {object} response.Response{data=safety.Swiper,msg=string} "查询成功"
// @Router /index_swiper/findSwiper [get]
func (index_swiperApi *SwiperApi) FindSwiper(c *gin.Context) {
	id := c.Query("id")
	reindex_swiper, err := index_swiperService.GetSwiper(id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reindex_swiper, c)
}

// GetSwiperList 分页获取轮播图列表
// @Tags Swiper
// @Summary 分页获取轮播图列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query safetyReq.SwiperSearch true "分页获取轮播图列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /index_swiper/getSwiperList [get]
func (index_swiperApi *SwiperApi) GetSwiperList(c *gin.Context) {
	var pageInfo safetyReq.SwiperSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := index_swiperService.GetSwiperInfoList(pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetSwiperPublic 不需要鉴权的轮播图接口
// @Tags Swiper
// @Summary 不需要鉴权的轮播图接口
// @accept application/json
// @Produce application/json
// @Param data query safetyReq.SwiperSearch true "分页获取轮播图列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /index_swiper/getSwiperPublic [get]
func (index_swiperApi *SwiperApi) GetSwiperPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    index_swiperService.GetSwiperPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的轮播图接口信息",
    }, "获取成功", c)
}
