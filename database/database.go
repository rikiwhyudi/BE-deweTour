package database

import (
	"BE-DeweTour/models"
	"BE-DeweTour/pkg/mysql"
	"fmt"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Country{},
		&models.Trip{},
		&models.Booking{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration failed!")
	}

	fmt.Println("Migration success")
}
