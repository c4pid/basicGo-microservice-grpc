package requests

import "time"

type CreateFlightRequest struct {
	Name          string    `json:"name" binding:"required"`
	From          string    `json:"from" binding:"required"`
	To            string    `json:"to" binding:"required"`
	Date          time.Time `json:"date" binding:"required"`
	Status        string    `json:"status" binding:"required"`
	AvailableSlot int64     `json:"available_slot" binding:"required"`
}

type UpdateFlightRequest struct {
	Id            int64     `json:"id" binding:"required"`
	Name          string    `json:"name" binding:"required"`
	From          string    `json:"from" binding:"required"`
	To            string    `json:"to" binding:"required"`
	Date          time.Time `json:"date" binding:"required"`
	Status        string    `json:"status" binding:"required"`
	AvailableSlot int64     `json:"available_slot" binding:"required"`
}

type SearchRequest struct {
	Name string `json:"name"`
	From string `json:"from"`
	To   string `json:"to"`
}
