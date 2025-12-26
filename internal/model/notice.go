package model

import "time"

type Notice struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Publisher string    `json:"publisher"` // 例如："物业中心"
	ViewCount int       `json:"view_count"`
	CreatedAt time.Time `json:"created_at"`
}

func (Notice) TableName() string {
	return "cms_notice"
}
