package models

import "gorm.io/gorm"

// Tweet represents a tweet in the system
type Tweet struct {
    gorm.Model
    UserID  uint   `json:"user_id"`
    Content string `json:"content"`
    User    User   `gorm:"foreignKey:UserID"`
}
