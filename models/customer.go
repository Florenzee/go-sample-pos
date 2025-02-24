package models

type Customer struct {
	CustomerID string `json:"customer_id" gorm:"primaryKey" column:"id"`
	Name       string `json:"name" gorm:"column:customer_name" column:"name; type: varchar(255);"`
	Email      string `json:"email" gorm:"column:customer_email" column:"email; type: varchar(255);"`
	Phone      string `json:"phone" gorm:"column:customer_phone" column:"phone; type: varchar(255);"`
	Address    string `json:"address"  gorm:"column:customer_address" column:"address; type: varchar(255);"`
	LoyaltyPts int    `json:"loyalty_points" gorm:"column:loyalty_pts; type: int(11);"`
}
