package request

type BookingRequest struct {
	RoomID       int    `json:"room_id"`
	CheckinDate  string `json:"checkin_date"`
	CheckoutDate string `json:"checkout_date"`
}
