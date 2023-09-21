package model

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null;unique"`
	Password  string    `json:"password" gorm:"not null"`
	RoleID    int       `json:"role_id" gorm:"not null;default:2"`
	Status    int       `json:"status" gorm:"not null;default:1"`
	UserID    *int      `json:"user_id" gorm:"default:null"`
	CreatedAt time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
