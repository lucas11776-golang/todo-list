package models

import "time"

type User struct {
	Connection string    `json:"-" connection:"default"`
	ID         int64     `json:"id" column:"id"`
	CreatedAt  time.Time `json:"created_at" column:"created_at"`
	FirstName  string    `json:"first_name" column:"first_name"`
	LastName   string    `json:"last_name" column:"last_name"`
	Email      string    `json:"email" column:"email"`
	Password   string    `json:"password" column:"password"`
}
