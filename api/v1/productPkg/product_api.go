package productPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type ProductApi struct {
}

func (m *ProductApi) CreateApi(c *gin.Context) {
	var product request.Product

	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := productApiService.CreateProduct(product.ProductName, product.ProductVendor, product.QuantityInStock, product.BuyPrice)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response.Ok(c)
}

func (m *ProductApi) DeleteApi(c *gin.Context) {
	var product request.Product
	err := c.ShouldBindJSON(&product)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = productApiService.DeleteProduct(product.Id)
	if err != nil {
		global.GVA_LOG.Error("刪除失敗!", zap.Error(err))
		response.FailWithMessage("刪除失敗", c)
		return
	}

	response.OkWithMessage("產品刪除成功", c)
}
