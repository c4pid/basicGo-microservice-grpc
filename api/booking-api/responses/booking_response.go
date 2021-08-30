package responses

import "time"

type BookingResponse struct {
	BookingCode int64     `json:"booking_code"`
	BookingDate time.Time `json:"booking_date"`
	CustomerId  int64     `json:"customer_id"`
	FlightId    int64     `json:"flight_id"`
	Status      string    `json:"status"`
}

type ViewBookingResponse struct {
	BookingCode int64                `json:"booking_code"`
	BookingDate time.Time            `json:"booking_date"`
	Customer    ViewCustomerResponse `json:"customer"`
	Flight      ViewFlightResponse   `json:"flight"`
}

type ViewCustomerResponse struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
}

type ViewFlightResponse struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	From          string `json:"from"`
	To            string `json:"to"`
	Date          time.Time
	Status        string `json:"status"`
	AvailableSlot int64  `json:"available_slot"`
}
