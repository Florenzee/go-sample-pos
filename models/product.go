package models

// atribut dari class atau struct
type Product struct {
	ProductID   string  `json:"product_id" gorm:"primaryKey" column:"id"`
	Name        string  `json:"name" gorm:"column:product_name" column:"name; type: varchar(255);"`
	Description string  `json:"description" gorm:"column:product_description" column:"description; type: varchar(255);"`
	Price       float64 `json:"price" gorm:"column:product_price" column:"price"`
	StockQty    int     `json:"stock_qty" gorm:"column:stock_qty" column:"stock_qty"`
	Category    string  `json:"category" gorm:"column:product_category" column:"category; type: varchar(255);"`
	SKU         string  `json:"sku" gorm:"column:product_sku" column:"sku; type: varchar(255);"`
	TaxRate     float64 `json:"tax_rate" gorm:"column:product_tax_rate" column:"tax_rate"`
}

type ProductError struct {
	Product Product `json:"product"`
	Error   error   `json:"error"`
}
