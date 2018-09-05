package main

type Selection struct {
	Id string `json:"id"`
	Input int `json:"input"`
	Seat int `json:"seat"`
	Occupied bool `json:"occupied"`
}

type Request struct {
	Rows int `json:"rows"`
	SeatsPerRow int `json:"seats_per_row"`
	Seats []Selection `json:"seats"`
}

