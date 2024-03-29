package Model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (b *User) TableName() string {
	return "user"
}
