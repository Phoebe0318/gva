package productPkg

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
)

type ProductService struct {
}

func (p *ProductService) CreateProduct(productName string, productVendor string, quantityInStock int, buyPrice float64) error {
	var count int64
	if err := global.GVA_DB.Model(&response.Product{}).Where("productName = ?", productName).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("產品名稱重複")
	}

	product := response.Product{
		ProductName:     productName,
		ProductVendor:   productVendor,
		QuantityInStock: quantityInStock,
		BuyPrice:        buyPrice,
	}

	if err := global.GVA_DB.Create(&product).Error; err != nil {
		return err
	}

	return nil

	//构建 SQL 插入语句
	//sql := "INSERT INTO products (productName, productVendor, quantityInStock, buyPrice) VALUES (?, ?, ?, ?)"
	//result := global.GVA_DB.Exec(sql, product.ProductName, product.ProductVendor, product.QuantityInStock, product.BuyPrice)
	//
	//// 检查是否发生错误
	//if result.Error != nil {
	//	return result.Error
	//}
	//
	//// 检查插入影响的行数
	//if result.RowsAffected == 0 {
	//	return errors.New("插入数据失败")
	//}
	//
	//return nil
}
