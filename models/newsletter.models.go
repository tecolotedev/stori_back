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
