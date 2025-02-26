package web

type CustomerCreateRequest struct {
	Name       string `validate:"required,min=1,max=100" json:"name"`
	Email      string `validate:"required,email" json:"email"`
	Phone      string `validate:"required,max=20" json:"phone"`
	Address    string `validate:"max=255" json:"address"`
	LoyaltyPts int    `validate:"min=0" json:"loyalty_points"`
}

type CustomerResponse struct {
	CustomerID string `json:"customer_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	LoyaltyPts int    `json:"loyalty_points"`
}

type CustomerUpdateRequest struct {
	CustomerID string `validate:"required" json:"customer_id"`
	Name       string `validate:"required,min=1,max=100" json:"name"`
	Email      string `validate:"required,email" json:"email"`
	Phone      string `validate:"required,max=20" json:"phone"`
	Address    string `validate:"max=255" json:"address"`
	LoyaltyPts int    `validate:"min=0" json:"loyalty_points"`
}
