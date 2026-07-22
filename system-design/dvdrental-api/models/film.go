package models

import (
	"time"

	"github.com/lib/pq"
)

type Film struct {
	FilmID          uint           `gorm:"column:film_id;primaryKey" json:"film_id"`
	Title           string         `gorm:"column:title" json:"title"`
	Description     string         `gorm:"column:description" json:"description"`
	ReleaseYear     int            `gorm:"column:release_year" json:"release_year"`
	LanguageID      uint           `gorm:"column:language_id" json:"language_id"`
	RentalDuration  int            `gorm:"column:rental_duration" json:"rental_duration"`
	RentalRate      float64        `gorm:"column:rental_rate" json:"rental_rate"`
	Length          int            `gorm:"column:length" json:"length"`
	ReplacementCost float64        `gorm:"column:replacement_cost" json:"replacement_cost"`
	Rating          string         `gorm:"column:rating" json:"rating"`
	LastUpdate      time.Time      `gorm:"column:last_update" json:"last_update"`
	SpecialFeatures pq.StringArray `gorm:"column:special_features;type:text[]" json:"special_features"`
}

func (Film) TableName() string {
	return "film"
}
