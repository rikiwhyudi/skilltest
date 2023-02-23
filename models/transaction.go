package models

import "time"

type Transaction struct {
	ID        int       `json:"id" gorm:"primary_key"`
	CartID    int       `json:"cart_id" gorm:"type:int"` //has many field carts
	UserID    int       `json:"user_id" gorm:"type:int"` //has many field users
	Total     int       `json:"total" gorm:"type:int"`
	CreatedAt time.Time `json:"created_at"`
}
