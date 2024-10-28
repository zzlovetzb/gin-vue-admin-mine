// 自动生成模板Swiper
package safety
import (
)

// 轮播图 结构体  Swiper
type Swiper struct {
    Id  *int `json:"id" form:"id" gorm:"primarykey;column:ID;comment:;"`  //序号 
    ImageURL  string `json:"www.baidu.com" form:"www.baidu.com" gorm:"column:imageURL;comment:;"`  //轮播图地址 
}


// TableName 轮播图 Swiper自定义表名 swiper
func (Swiper) TableName() string {
    return "swiper"
}

