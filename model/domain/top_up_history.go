package domain

import "time"

type TopUpHistory struct {
	ID            int       `json:"id"`
	WalletID      int       `json:"wallet_id"`
	AmountTopUp   float64   `json:"amount_top_up"`
	TopUpDateTime time.Time `json:"top_up_date_time"`
}
