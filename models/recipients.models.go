package models

import "time"

type Recipient struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Newsletters []Newsletter `json:"newsletters" gorm:"many2many:recipient_newsletter;"`
}
