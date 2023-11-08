package time

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ApiGroup struct {
}

func GetCurrentTime(c *gin.Context) {
	var request struct {
		Message string `json:"message"`
	}

	// 从 POST 请求的 JSON 数据中解析参数
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 检查输入参数
	if request.Message != "current time" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid message.",
		})
		return
	}

	// 获取当前时间
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	// 返回当前时间
	c.JSON(http.StatusOK, gin.H{
		"current_time": currentTime,
	})
}
