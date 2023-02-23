package database

import (
	"e-market/models"
	"e-market/pkg/mysql"
	"fmt"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.CartItem{},
		&models.Cart{},
		&models.Transaction{},
	)

	if err != nil {
		panic("failed to migrated")
	}

	fmt.Printf("successfully migrated")
}
