package response

type Seat struct {
	Row    uint `json:"row"`
	Column uint `json:"column"`
}

type BookingSeats struct {
	Base
	Seats []Seat `json:"seats"`
}

type GetAvailableSeats struct {
	Base
	AvailableSeats []Seat `json:"available_seats"`
}
