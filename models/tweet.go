package models

import "gorm.io/gorm"

// Tweet represents a tweet in the system
type Tweet struct {
    gorm.Model
    Content string `gorm:"not null"`
    UserID  uint   // Foreign key to User
}
