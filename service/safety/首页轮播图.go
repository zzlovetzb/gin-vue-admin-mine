package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
    safetyReq "github.com/flipped-aurora/gin-vue-admin/server/model/safety/request"
)

type SwiperService struct {}
// CreateSwiper 创建轮播图记录
// Author [piexlmax](https://github.com/piexlmax)
func (index_swiperService *SwiperService) CreateSwiper(index_swiper *safety.Swiper) (err error) {
	err = global.GVA_DB.Create(index_swiper).Error
	return err
}

// DeleteSwiper 删除轮播图记录
// Author [piexlmax](https://github.com/piexlmax)
func (index_swiperService *SwiperService)DeleteSwiper(id string) (err error) {
	err = global.GVA_DB.Delete(&safety.Swiper{},"ID = ?",id).Error
	return err
}

// DeleteSwiperByIds 批量删除轮播图记录
// Author [piexlmax](https://github.com/piexlmax)
func (index_swiperService *SwiperService)DeleteSwiperByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]safety.Swiper{},"ID in ?",ids).Error
	return err
}

// UpdateSwiper 更新轮播图记录
// Author [piexlmax](https://github.com/piexlmax)
func (index_swiperService *SwiperService)UpdateSwiper(index_swiper safety.Swiper) (err error) {
	err = global.GVA_DB.Model(&safety.Swiper{}).Where("ID = ?",index_swiper.Id).Updates(&index_swiper).Error
	return err
}

// GetSwiper 根据id获取轮播图记录
// Author [piexlmax](https://github.com/piexlmax)
func (index_swiperService *SwiperService)GetSwiper(id string) (index_swiper safety.Swiper, err error) {
	err = global.GVA_DB.Where("ID = ?", id).First(&index_swiper).Error
	return
}

// GetSwiperInfoList 分页获取轮播图记录
// Author [piexlmax](https://github.com/piexlmax)
func (index_swiperService *SwiperService)GetSwiperInfoList(info safetyReq.SwiperSearch) (list []safety.Swiper, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&safety.Swiper{})
    var index_swipers []safety.Swiper
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&index_swipers).Error
	return  index_swipers, total, err
}
func (index_swiperService *SwiperService)GetSwiperPublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
