package response

type Room struct {
	ID            int    `json:"id"`
	Row           int    `json:"row"`
	Column        int    `json:"column"`
	MinDistance   int    `json:"min_distance"`
	AvailableSeat []Seat `json:"available_seat"`
	BookedSeat    []Seat `json:"booked_seat"`
}

type GetRoomDetail struct {
	Base
	Data *Room `json:"data"`
}

type GetRoomList struct {
	Base
	Data  []*Room `json:"data"`
	Total int     `json:"total"`
}

type CreateRoom struct {
	Base
	Data *Room `json:"data"`
}

type UpdateRoom struct {
	Base
	Data *Room `json:"data"`
}

type DeleteRoom struct {
	Base
}
