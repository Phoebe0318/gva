package myPkg

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type MyApi struct {
}

func (e *MyApi) InitMyApiRouter(Router *gin.RouterGroup) {
	MyRouterGroup := Router.Group("myapi")
	api := v1.ApiGroupApp.MyPkgApiGroup.MyApi
	{
		MyRouterGroup.POST("create", api.CreateApi)
	}
}
