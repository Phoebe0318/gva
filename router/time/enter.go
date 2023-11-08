package time

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/time"
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
}

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 创建路由组
	api := r.Group("/api")

	// 注册 GetCurrentTime 接口
	api.POST("/getcurrenttime", time.GetCurrentTime)

	return r
}
