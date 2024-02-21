package images

import (
	"time"

	"gorm.io/gorm"
)

type Image struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"username"`
	Category    string         `json:"email"`
	Description string         `json:"description"`
	File        string         `json:"file"`
	Tags        string         `json:"tags"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeleteAt    gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}
