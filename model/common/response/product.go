package response

type Product struct {
	ProductName     string  `gorm:"column:productName;type:varchar(70);comment:產品名稱" json:"productName"`
	ProductVendor   string  `gorm:"column:productVendor;type:varchar(50);comment:供應商" json:"productVendor"`
	QuantityInStock int     `gorm:"column:quantityInStock;type:smallint;comment:庫存數量" json:"quantityInStock"`
	BuyPrice        float64 `gorm:"column:buyPrice;type:decimal(10,2);comment:購買價格" json:"buyPrice"`
}

func (Product) TableName() string {
	return "products"
}
