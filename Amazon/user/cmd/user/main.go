package main

import (
	"Amazon/user/internal/models"
	"fmt"
)

func main() {
	user := models.User{
		ID:           "1",
		FirstName:    "amol",
		LastName:     "gajare",
		Email:        "gajre425@gmail.com",
		PasswordHash: "today!098",
		Address: []models.Address{
			models.Address{
				ID:      "100",
				UserID:  "gajare",
				Street:  "praksh uday",
				City:    "pune",
				State:   "maharstra",
				PinCode: "41106",
			},models.Address{
				ID:      "100",
				UserID:  "gajare",
				Street:  "praksh uday",
				City:    "pune",
				State:   "maharstra",
				PinCode: "41106",
			},
		},
	}
	fmt.Println("hello :", user)
}
