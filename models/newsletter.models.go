package models

import (
	"time"
)

type Newsletter struct {
	ID                 uint                `gorm:"primarykey" json:"id"`
	Name               string              `json:"name"`
	NewsletterVersions []NewsletterVersion `json:"newsletter_versions"`
	CreatedAt          time.Time           `json:"created_at"`
	UpdatedAt          time.Time           `json:"updated_at"`
}

type NewsletterVersion struct {
	ID           uint      `gorm:"primarykey" json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	File         string    `json:"file"`
	NewsletterID uint      `json:"newsletter_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
