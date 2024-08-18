package models

import "github.com/jinzhu/gorm"

type User struct {
    gorm.Model
    Username string `gorm:"unique;not null" json:"username"`
    Password string `gorm:"not null" json:"password"`
    Email    string `gorm:"unique;not null" json:"email"`
    Role     string `gorm:"default:'Regular'" json:"role"`
}
