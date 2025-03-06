package response

type Seat struct {
	Row    uint `json:"row" example:"0"`
	Column uint `json:"column" example:"0"`
}

type BookingSeats struct {
	Base
	Seats []Seat `json:"seats"`
}

type GetAvailableSeats struct {
	Base
	AvailableSeats []Seat `json:"available_seats"`
}
