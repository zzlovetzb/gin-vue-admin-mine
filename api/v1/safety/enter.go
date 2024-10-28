package safety

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ SwiperApi }

var index_swiperService = service.ServiceGroupApp.SafetyServiceGroup.SwiperService
