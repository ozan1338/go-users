package main

import (
	"go-users/src/database"
	"go-users/src/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	database.Connect()
	db, err := gorm.Open(mysql.Open("root:root@tcp(host.docker.internal:33066)/ambassador"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	var users []models.User

	db.Find(&users)

	for _, user := range users {
		database.DB.Create(&user)
	}



	// for i := 0; i < 30; i++ {
	// 	ambassador := models.User{
	// 		FirstName:    faker.FirstName(),
	// 		LastName:     faker.LastName(),
	// 		Email:        faker.Email(),
	// 		IsAmbassador: true,
	// 	}

	// 	ambassador.SetPassword("1234")

	// 	database.DB.Create(&ambassador)
	// }
}
