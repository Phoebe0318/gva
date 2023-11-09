package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/myPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/productPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/time"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	TimeApiGroup    time.ApiGroup
	MyPkgApiGroup   myPkg.ApiGroup
	ProductApiGroup productPkg.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
