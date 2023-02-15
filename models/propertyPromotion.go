package models

type PropertyPromotion struct {
	Id uint64 `gorm:"primaryKey; not null" json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Image_Banner string `json:"image_banner"`
	Start_Time string `json:"start_time"`
	End_Time string `json:"end_time"`
}

type PropertyPromotionForm struct {
	Title string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Image_Banner string `json:"image_banner" binding:"required"`
	Start_Time string `json:"start_time" binding:"required"`
	End_Time string `json:"end_time" binding:"required"`
}