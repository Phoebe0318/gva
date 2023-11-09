package productPkg

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	ProductApi
}

var (
	productApiService = service.AppServices.ProductServiceGroup.ProductService
)
