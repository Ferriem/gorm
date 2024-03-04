package main

import (
	"fmt"

	"github.com/Ferriem/gorm/code/common"
	"github.com/Ferriem/gorm/code/datamodels"
	"github.com/Ferriem/gorm/code/repositories"
	"github.com/Ferriem/gorm/code/services"
)

func failOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := common.NewConn()
	failOnErr(err)
	fmt.Println("Connection to MySQL was successful")

	err = db.AutoMigrate(&datamodels.User{}, &datamodels.CreditCard{})
	failOnErr(err)
	fmt.Println("Auto migration was successful")

	user := &datamodels.User{
		Name:       "Ferriem",
		Email:      nil,
		Age:        20,
		CreditCard: datamodels.CreditCard{Number: "777777"},
	}
	userRepository := repositories.NewUserManager("users", db)
	userService := services.NewUserService(userRepository)
	result, err := userService.InsertUser(user)
	failOnErr(err)
	fmt.Printf("User %d was created successfully\n", result)

	users := []*datamodels.User{
		&datamodels.User{Name: "John", Age: 18},
		&datamodels.User{Name: "Jack", Age: 18},
	}
	rowsAffected, err := userService.InsertUsers(users)
	failOnErr(err)
	fmt.Printf("%d users was created successfully\n", rowsAffected)

	data := map[string]interface{}{
		"Name": "John Doe",
		"Age":  30,
	}
	ans := db.Model(&datamodels.User{}).Create(data)
	failOnErr(err)
	fmt.Printf("%d users was created successfully\n", ans.RowsAffected)

	fmt.Printf("All User Name:\n")
	users, err = userService.GetAll()
	failOnErr(err)
	for i := range users {
		fmt.Println(users[i].Name)
	}

	user, err = userService.GetUserByID(1)
	failOnErr(err)
	fmt.Printf("Before Update UserID: %d, Name: %s, CreditCard: %s\n", user.ID, user.Name, user.CreditCard.Number)

	user.CreditCard.Number = "0427"

	err = userService.UpdateUser(user)
	failOnErr(err)
	fmt.Printf("Successfully update\n")

	user, err = userService.GetUserByID(1)
	failOnErr(err)
	fmt.Printf("After Update UserID: %d, Name: %s, CreditCard: %s\n", user.ID, user.Name, user.CreditCard.Number)

	err = userService.DeleteUser(1)
	failOnErr(err)
	fmt.Printf("Successfully delete\n")

	fmt.Printf("After delete:\n")
	users, err = userService.GetAll()
	failOnErr(err)
	for i := range users {
		fmt.Println(users[i].Name)
	}

}
