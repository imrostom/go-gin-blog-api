package models

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID              uint       `gorm:"primarykey" json:"id"`
	Name            string     `gorm:"type:varchar(100);not null" json:"name"`         // User's full name
	Image           string     `gorm:"type:varchar(255);null" json:"image"`            // Optional profile picture URL
	Email           string     `gorm:"type:varchar(100);unique;not null" json:"email"` // Email address, must be unique
	Password        string     `gorm:"type:varchar(255);not null" json:"-"`            // Hashed password, hidden in API responses
	Phone           string     `gorm:"type:varchar(20);null" json:"phone"`             // Optional phone number
	DateOfBirth     time.Time  `gorm:"type:date;null" json:"date_of_birth"`            // Optional date of birth
	Address         string     `gorm:"type:varchar(255);null" json:"address"`          // Optional address
	Status          uint8      `gorm:"default:1" json:"status"`                        // Account status: 1 = active, 0 = suspended, etc.
	EmailVerifiedAt *time.Time `gorm:"null" json:"email_verified_at"`                  // Timestamp of email verification
	LastLoginAt     *time.Time `gorm:"null" json:"last_login_at"`                      // Timestamp of the user's last login
	Role            string     `gorm:"type:varchar(50);default:'user'" json:"role"`    // Role (e.g., user, admin)
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}
