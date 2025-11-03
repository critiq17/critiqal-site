package user

import "gorm.io/gorm"

type User struct {
	ID        string
	Username  string
	Email     string
	Password  string
	FirstName string
	LastName  string
	PhotoURL  string
	CreatedAt int64
	DeletedAt gorm.DeletedAt
}
