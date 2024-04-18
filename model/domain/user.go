package domain

type User struct {
	ID       int
	FullName string
	Email    string
	Password string
	Age      int
	Role     string
	Wallet   Wallet `json:"wallet" gorm:"foreignKey:UserID;references:ID"`
}
