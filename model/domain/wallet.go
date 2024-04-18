package domain

type Wallet struct {
	ID           int            `json:"id"`
	UserID       int            `json:"user_id"`
	Balance      float64        `json:"balance"`
	TopUpHistory []TopUpHistory `json:"top_up_history,omitempty" gorm:"foreignKey:WalletID;references:ID"`
}
