package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserDto
	CommonModel
}

type UserDto struct {
	FirstName  string `json: "first_name"`
	MiddleName string `json: "middle_name"`
	LastName   string `json: "last_name"`
	Age        uint16 `json: "age"`
	Email      string `json: "email"`
	Password   string `json: "password"`
	Address    string `json: "address"`
	
}

func FetAllUsers() []User {

	user1 := User{
		Model: gorm.Model{
			ID: 1,
		},
		CommonModel: CommonModel{Status: true, Priority: 1},
		UserDto: UserDto{
			Email:     "sabin@gmail.com",
			FirstName: "sabin",
			LastName:  "Kc",
			Age:       22,
			Address:   "Sindhuli",
			Password:  "password",
		},
	}
	return []User{
		user1,
	}
}
