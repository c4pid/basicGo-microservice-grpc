package responses

import "time"

type FlightResponse struct {
	Id            int64     `json:"id"`
	Name          string    `json:"name"`
	From          string    `json:"from"`
	To            string    `json:"to"`
	Date          time.Time `json:"date"`
	Status        string    `json:"status"`
	AvailableSlot int64     `json:"available_slot"`
}

type ListFlights struct {
	Flights []*FlightResponse `json:"flights"`
}
