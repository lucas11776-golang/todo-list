package models

import "time"

type Task struct {
	Connection  string    `json:"-" connection:"sqlite" table:"tasks"`
	ID          int64     `json:"id" column:"id" type:"primary_key"`
	UserID      int64     `json:"-" column:"user_id" type:"integer"`
	CreatedAt   time.Time `json:"created_at" column:"created_at" type:"datetime_current"`
	DueDate     time.Time `json:"due_date" column:"due_date" type:"datetime"`
	Title       string    `json:"title" column:"title" type:"string"`
	Description string    `json:"description" column:"description" type:"text"`
	Complete    bool      `json:"complete" column:"complete" type:"boolean"`
}
