package models

import (
	"time"
)

// Post represents a blog post or article
type Post struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`       // Post title
	Slug        string    `gorm:"type:varchar(255);unique;not null" json:"slug"` // URL-friendly version of the title
	UserId      uint      `gorm:"not null" json:"user_id"`                       // Reference to the user who created the post
	CategoryId  uint      `gorm:"not null" json:"category_id"`                   // Reference to the post category
	Content     string    `gorm:"type:text;not null" json:"content"`             // Post content
	Image       string    `gorm:"type:varchar(255); null" json:"image"`          // URL or path to the image
	Status      uint8     `gorm:"default:1" json:"status"`                       // Post status (active/inactive)
	Views       uint      `gorm:"default:0" json:"views"`                        // Number of views
	PublishedAt time.Time `json:"published_at"`                                  // When the post was published
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
