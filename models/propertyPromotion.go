package models

import "time"

type PropertyPromotion struct {
	Id uint64 `gorm:"primaryKey; not null" json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	ImageBanner string `json:"image_banner"`
	StartTime time.Time `json:"start_time"`
	EndTime time.Time `json:"end_time"`
}