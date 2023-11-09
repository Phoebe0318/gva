package productPkg

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type ProductApi struct {
}

func (e *ProductApi) InitProductApiRouter(Router *gin.RouterGroup) {
	ProductRouterGroup := Router.Group("product")
	api := v1.ApiGroupApp.ProductApiGroup.ProductApi
	{
		ProductRouterGroup.POST("create", api.CreateApi)
	}
}
