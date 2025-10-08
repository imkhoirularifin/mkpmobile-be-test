package dto

import "time"

type ScheduleDto struct {
	ID             uint      `json:"id"`
	MovieTitle     string    `json:"movie_title"`
	StudioName     string    `json:"studio_name"`
	ShowDate       time.Time `json:"show_date"`
	ShowTime       string    `json:"show_time"`
	AvailableSeats int       `json:"available_seats"`
	Price          float64   `json:"price"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type CreateScheduleRequest struct {
	MovieTitle     string  `json:"movie_title" validate:"required,min=1,max=255"`
	StudioName     string  `json:"studio_name" validate:"required,min=1,max=100"`
	ShowDate       string  `json:"show_date" validate:"required,datetime=2006-01-02"`
	ShowTime       string  `json:"show_time" validate:"required,datetime=15:04"`
	AvailableSeats int     `json:"available_seats" validate:"required,min=0"`
	Price          float64 `json:"price" validate:"required,min=0"`
}

type UpdateScheduleRequest struct {
	MovieTitle     *string  `json:"movie_title" validate:"omitempty,min=1,max=255"`
	StudioName     *string  `json:"studio_name" validate:"omitempty,min=1,max=100"`
	ShowDate       *string  `json:"show_date" validate:"omitempty,datetime=2006-01-02"`
	ShowTime       *string  `json:"show_time" validate:"omitempty,datetime=15:04"`
	AvailableSeats *int     `json:"available_seats" validate:"omitempty,min=0"`
	Price          *float64 `json:"price" validate:"omitempty,min=0"`
}
