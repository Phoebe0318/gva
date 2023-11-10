package request

type Product struct {
	Id              int     `gorm:"column:id;type:int(11)" json:"id"`
	ProductName     string  `gorm:"column:productName;type:varchar(70)" json:"productName"`
	ProductVendor   string  `gorm:"column:productVendor;type:varchar(50)" json:"productVendor"`
	QuantityInStock int     `gorm:"column:quantityInStock;type:smallint" json:"quantityInStock"`
	BuyPrice        float64 `gorm:"column:buyPrice;type:decimal(10,2)" json:"buyPrice"`
}

func (Product) TableName() string {
	return "products"
}
