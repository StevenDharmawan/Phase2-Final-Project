package request

type RoomRequest struct {
	RoomNumber   string  `json:"room_number"`
	Availability bool    `json:"availability"`
	Price        float64 `json:"price"`
	Category     string  `json:"category"`
}
