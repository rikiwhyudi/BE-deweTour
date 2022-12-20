package models

import "time"

type Country struct {
	ID   int    `json:"id"`
	Name string `json:"name" gorm:"type: varchar(255)"`
}

type Trip struct {
	ID             int       `json:"id" gorm:"primary_key:auto_increment"`
	Title          string    `json:"title" gorm:"type: varchar(255)"`
	CountryID      int       `json:"-"`
	Country        Country   `json:"country" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Accomodation   string    `json:"accomodation" gorm:"type: varchar(255)"`
	Transportation string    `json:"transportation" gorm:"type: varchar(255)"`
	Eat            string    `json:"eat" gorm:"type: varchar(255)"`
	Day            int       `json:"day" gorm:"type: varchar(50)"`
	Night          int       `json:"night" gorm:"type: varchar(50)"`
	DateTrip       time.Time `json:"dateTrip" gorm:"type: varchar(255)"`
	Price          int       `json:"price" gorm:"type: varchar(255)"`
	Quota          int       `json:"quota" gorm:"type: varchar(255)"`
	Description    string    `json:"description" gorm:"type: varchar(255)"`
	Image          string    `json:"image" gorm:"type: varchar(255)"`
}
