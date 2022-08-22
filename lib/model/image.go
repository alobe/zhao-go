package model

type Image struct {
	ID        int    `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt int64  `gorm:"autoCreateTime:milli" json:"created_at"`
	Url       string `gorm:"unique" json:"url"`
}
