package response

type Room struct {
	ID            int    `json:"id" example:"1"`
	Row           int    `json:"row" example:"4"`
	Column        int    `json:"column" example:"5"`
	MinDistance   int    `json:"min_distance" example:"6"`
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
