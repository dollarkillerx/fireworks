package base

import (
	"github.com/dollarkillerx/fireworks/internal/pkg/enum"
	"github.com/dollarkillerx/fireworks/internal/pkg/models"

	"log"
)

func (b *Base) CreateUser(email string, name string, password string, role enum.Role) error {
	err := b.db.Model(&models.User{}).Create(&models.User{
		Email:    email,
		UserName: name,
		Password: password,
		Role:     role,
	}).Error
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (b *Base) GetUser(email string) (*models.User, error) {
	var us models.User
	err := b.db.Model(&models.User{}).Where("email = ?", email).First(&us).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &us, nil
}
