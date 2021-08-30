package requests

import "time"

type CreateBooking struct {
	BookingDate time.Time `json:"booking_date" binding:"required"`
	CustomerId  int64     `json:"customer_id" binding:"required"`
	FlightId    int64     `json:"flight_id" binding:"required"`
	Status      string    `json:"status" binding:"required"`
}

type ViewBooking struct {
	BookingCode int64 `json:"booking_code"`
}
