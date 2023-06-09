package admin

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}

func (repo Repository) CheckUsername(username, password string) (*User, error) {
	var user User
	result := repo.DB.Where("username", username).Where("password", password).First(&user)

	return &user, result.Error
}