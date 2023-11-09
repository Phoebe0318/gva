package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/myPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/router/productPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/time"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Time    time.RouterGroup
	MyPkg   myPkg.RouterGroup
	Product productPkg.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
