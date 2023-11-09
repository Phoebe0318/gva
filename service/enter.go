package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/myPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/service/productPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/time"
)

type AppServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	TimeServiceGroup    time.ServiceGroup
	MyPkgServiceGroup   myPkg.ServiceGroup
	ProductServiceGroup productPkg.ServiceGroup
}

var AppServices = new(AppServiceGroup)
