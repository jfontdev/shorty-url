package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Url struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	OriginalURL  string    	`gorm:"not null"`
	ShortenedURL string    	`gorm:"not null"`
	gorm.Model
}

type Urls struct {
	Urls []Url `json:"urls"`
}

func (url *Url) BeforeCreate(tx *gorm.DB) (err error) {
	url.ID = uuid.New()
	return
}
