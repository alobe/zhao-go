package model

type Image struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	CreatedAt int64  `gorm:"autoCreateTime" json:"created_at"`
	Url       string `gorm:"unique" json:"url"`
}
