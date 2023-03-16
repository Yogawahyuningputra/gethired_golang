package database

import (
	"backend/models"
	"backend/packages/mysql"
	"fmt"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.Activity{},
		&models.Todos{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Successfull !!")
}
