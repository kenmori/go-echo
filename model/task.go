package model

import "time"

type Task struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Title     string    `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
	User      User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId    uint      `json:"user_id" gorm:"not null"` // 1対多の関係を表す外部キー
}

type TaskResponse struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Title     string    `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}
