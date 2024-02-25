package model

import "time"

type Todo struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Content   string    `json:"content" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	Checked   bool      `json:"checked"`
	User      User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId    uint      `json:"user_id" gorm:"not null"`
}

type TodoResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Content   string    `json:"content" gorm:"not null"`
	Checked   bool      `json:"checked"`
	CreatedAt time.Time `json:"created_at"`
}
