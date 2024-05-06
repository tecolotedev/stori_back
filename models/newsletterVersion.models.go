package models

import "time"

type NewsletterVersion struct {
	ID           uint      `gorm:"primarykey" json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	File         string    `json:"file"`
	Sent         bool      `json:"sent" gorm:"default:false"`
	NewsletterID uint      `json:"newsletter_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
