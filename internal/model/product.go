package model

import "time"

type Product struct {
	ID            int64     `gorm:"primaryKey" json:"id"`
	CategoryName  string    `json:"category_name"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Price         float64   `json:"price"`
	OriginalPrice float64   `json:"original_price"`
	Stock         int       `json:"stock"`
	ImageUrl      string    `json:"image_url"`
	IsPromotion   int       `json:"is_promotion"` // 1:是 0:否
	Sales         int       `json:"sales"`
	Status        int       `json:"status"` // 1:上架 0:下架
	CreatedAt     time.Time `json:"created_at"`
}

func (Product) TableName() string {
	return "pms_product"
}
