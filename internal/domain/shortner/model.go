package shortner

import "time"

type Shortner struct {
	ID            uint      `gorm:"primarykey"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Url           string    `json:"url" gorm:"type:text;not null;index"`
	Code          string    `json:"code" gorm:"type:varchar(10);uniqueIndex;not null"`
	Redirect      bool      `json:"redirect" gorm:"default:false"`
	RedirectCount int       `json:"redirect_count" gorm:"default:0"`
}
