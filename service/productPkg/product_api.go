package productPkg

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ProductService struct {
}

func (p *ProductService) CreateProduct(productName string, productVendor string, quantityInStock int, buyPrice float64) error {
	var count int64
	if err := global.GVA_DB.Model(&request.Product{}).Where("productName = ?", productName).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("產品名稱重複")
	}

	product := request.Product{
		ProductName:     productName,
		ProductVendor:   productVendor,
		QuantityInStock: quantityInStock,
		BuyPrice:        buyPrice,
	}

	if err := global.GVA_DB.Create(&product).Error; err != nil {
		return err
	}

	return nil
}

func (p *ProductService) DeleteProduct(id int) error {
	var count int64
	if err := global.GVA_DB.Model(&request.Product{}).Where("id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return errors.New("產品不存在")
	}

	if err := global.GVA_DB.Where("id = ?", id).Delete(&request.Product{}).Error; err != nil {
		return err
	}

	return nil
}
