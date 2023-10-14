package db

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	gorm.Model
	Code  string `json:"code"`
	Price uint   `json:"price"`
}

type Role struct {
	ID        uint       `json:"id"`
	RoleName  string     `json:"role_name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type User struct {
}
