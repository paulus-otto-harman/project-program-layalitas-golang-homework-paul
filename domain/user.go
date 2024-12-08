package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `gorm:"type:varchar(20);not null"`
	Email string `gorm:"type:varchar(20);not null"`
}

func UserSeed() []User {
	return []User{
		{
			Name:  "User Satu",
			Email: "user@mail.com",
		},
	}
}
