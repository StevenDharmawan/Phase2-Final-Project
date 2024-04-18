package response

import "phase2-final-project/model/domain"

type UserResponse struct {
	ID       int           `json:"id"`
	FullName string        `json:"full_name"`
	Email    string        `json:"email"`
	Age      int           `json:"age"`
	Wallet   domain.Wallet `json:"wallet,omitempty" gorm:"foreignKey:UserID;references:ID"`
}
