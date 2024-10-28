package safety

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ SwiperRouter }

var index_swiperApi = api.ApiGroupApp.SafetyApiGroup.SwiperApi
