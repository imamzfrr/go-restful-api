package domain

type Product struct {
	ProductID   string  `gorm:"primaryKey;column:product_id" json:"product_id"`
	Name        string  `gorm:"column:name" json:"name"`
	Description string  `gorm:"column:description" json:"description"`
	Price       float64 `gorm:"column:price" json:"price"`
	StockQty    int     `gorm:"column:stock_qty" json:"stock_qty"`
	Category    string  `gorm:"column:category" json:"category"`
	SKU         string  `gorm:"column:sku" json:"sku"`
	TaxRate     float64 `gorm:"column:tax_rate" json:"tax_rate"`
}
