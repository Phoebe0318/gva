package productPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductApi struct {
}

func (m *ProductApi) CreateApi(c *gin.Context) {
	var product response.Product

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
