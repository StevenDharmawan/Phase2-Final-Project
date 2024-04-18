package domain

import "time"

type Booking struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	RoomID       int       `json:"room_id"`
	CheckinDate  time.Time `json:"checkin_date"`
	CheckoutDate time.Time `json:"checkout_date"`
	Status       string    `json:"status"`
	TotalPrice   float64   `json:"total_price"`
}
