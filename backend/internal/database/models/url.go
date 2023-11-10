package models

import (
	"time"
)

type Url struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoincrement"`
	LongUrl   string    `json:"fqdn" gorm:"type:varchar(255)"`
	ShortUrl  string    `json:"short_url" gorm:"type:varchar(255)"`
	ShortCode string    `json:"short_code" gorm:"type:varchar(255)"`
	ExpireAt  time.Time `json:"expire_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
