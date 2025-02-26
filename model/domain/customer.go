package domain

type Customer struct {
	CustomerID string `gorm:"primaryKey;column:customer_id" json:"customer_id"`
	Name       string `gorm:"column:name" json:"name"`
	Email      string `gorm:"column:email" json:"email"`
	Phone      string `gorm:"column:phone" json:"phone"`
	Address    string `gorm:"column:address" json:"address"`
	LoyaltyPts int    `gorm:"column:loyalty_points" json:"loyalty_points"`
}
