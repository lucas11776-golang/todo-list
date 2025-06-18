package models

import "time"

type User struct {
	Connection string    `json:"-" connection:"sqlite"`
	ID         int64     `json:"id" column:"id" type:"primary_key"`
	CreatedAt  time.Time `json:"created_at" column:"created_at" type:"datetime_current"`
	FirstName  string    `json:"first_name" column:"first_name" type:"string"`
	LastName   string    `json:"last_name" column:"last_name" type:"string"`
	Email      string    `json:"email" column:"email" type:"string"`
	Password   string    `json:"password" column:"password" type:"string"`
}
