package models

import "time"

type Product struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" gorm:"type:varchar(255)"`
	Price     int       `json:"price" gorm:"type:int"`
	Stock     int       `json:"stock" gorm:"type:int"`
	Cart      []Cart    `json:"cart" gorm:"many2many:cart_items;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
