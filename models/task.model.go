package models

import "time"

type Task struct {
	Connection  string    `json:"-" connection:"default" table:"tasks"`
	ID          int64     `json:"id" column:"id"`
	UserID      int64     `json:"-" column:"user_id"`
	CreatedAt   time.Time `json:"created_at" column:"created_at"`
	Title       string    `json:"title" column:"title"`
	Description string    `json:"description" column:"description"`
	Complete    bool      `json:"complete" column:"complete"`
}
