package entity

import (
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	MovieTitle     string         `gorm:"type:varchar(255);not null" json:"movie_title"`
	StudioName     string         `gorm:"type:varchar(100);not null" json:"studio_name"`
	ShowDate       time.Time      `gorm:"type:date;not null" json:"show_date"`
	ShowTime       string         `gorm:"type:varchar(10);not null" json:"show_time"`
	AvailableSeats int            `gorm:"not null" json:"available_seats"`
	Price          float64        `gorm:"type:decimal(10,2);not null" json:"price"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
