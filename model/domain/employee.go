package domain

type Employee struct {
	EmployeeID string `gorm:"primaryKey;column:employee_id" json:"employee_id"`
	Name       string `gorm:"column:name" json:"name"`
	Role       string `gorm:"column:role" json:"role"`
	Email      string `gorm:"column:email" json:"email"`
	Phone      string `gorm:"column:phone" json:"phone"`
	DateHired  string `gorm:"column:date_hired" json:"date_hired"`
}
