package models

import "time"

type Booking struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	UserID    int       `json:"user_id"`
	User      User      `json:"user"`
	TripID    int       `json:"trip_id"`
	Trip      Trip      `json:"trip" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Qty       int       `json:"qty"`
	SubTotal  int       `json:"sub_total" gorm:"type:int"`
	UpdatedAt time.Time `json:"updated_at"`
}
