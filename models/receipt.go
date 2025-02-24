package models

type Receipt struct {
	ReceiptID   string  `json:"receipt_id" gorm:"primary_key" column:"receipt_id"`
	OrderID     string  `json:"order_id" gorm:"column:order_id" column:"order_id"`
	PaymentID   string  `json:"payment_id" gorm:"column:payment_id" column:"payment_id"`
	ReceiptDate string  `json:"receipt_date" gorm:"column:receipt_date" column:"receipt_date"`
	TotalAmount float64 `json:"total_amount" gorm:"column:total_amount" column:"total_amount"`
	Taxes       float64 `json:"taxes" gorm:"column:taxes" column:"taxes"`
	Discount    float64 `json:"discount" gorm:"column:discount" column:"discount"`
	FinalAmount float64 `json:"final_amount" gorm:"column:final_amount" column:"final_amount"`
}
