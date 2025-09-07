package models

import "gorm.io/gorm"

// 数据模型
type ShortLink struct {
	gorm.Model
	ShortCode   string `gorm:"column:short_code;uniqueIndex"`
	OriginalURL string `gorm:"column:original_url"`
	Password    string `gorm:"column:password"`
	ExpiresAt   int64  `gorm:"column:expires_at"`
}
