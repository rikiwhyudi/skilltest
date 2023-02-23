package models

type Cart struct {
	ID           int           `json:"id" gorm:"primary_key"`
	UserID       int           `json:"user_id" gorm:"type:int"`
	Items        []Product     `json:"items" gorm:"many2many:cart_items"`
	Transactions []Transaction `json:"transactions" gorm:"foreignKey:CartID"`
}

type CartItem struct {
	ProductID int     `json:"product_id" gorm:"type:int"`
	Product   Product `json:"product"`
	CartID    int     `json:"cart_id" gorm:"type:int"`
	Cart      Cart    `json:"cart"`
	Subtotal  int     `json:"subtotal" gorm:"type:int"`
	Quantity  int     `json:"quantity" gorm:"type:int"`
}
