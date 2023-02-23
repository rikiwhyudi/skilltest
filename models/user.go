package models

type User struct {
	ID           int           `json:"id" gorm:"primary_key"`
	Name         string        `json:"name" gorm:"type:varchar(255)"`
	Email        string        `json:"email" gorm:"type:varchar(255)"`
	Password     string        `json:"-" gorm:"type:varchar(255)"`
	Transactions []Transaction `json:"transactions" gorm:"foreignKey:UserID"` //has many transactions
	Carts        []Cart        `json:"carts" gorm:"foreignKey:UserID"`        //has many carts
}
