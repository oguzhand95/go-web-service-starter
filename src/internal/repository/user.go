package repository

import (
	"github.com/oguzhand95/go-web-service-starter/src/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	Database *gorm.DB
}

func NewUserRepository(database *Database) (*UserRepository, error) {
	err := database.Gorm.AutoMigrate(&model.User{})

	if err != nil {
		return nil, err
	}

	return &UserRepository{
		Database: database.Gorm,
	}, nil
}

func (lr *UserRepository) Register(user *model.User) error {
	if db := lr.Database.Create(&user); db.Error != nil {
		return db.Error
	}

	return nil
}

func (lr *UserRepository) GetByFieldMail(mail string) (*model.User, error) {
	var user model.User

	db := lr.Database.Where("mail = ?", mail).First(&user)

	if db.Error != nil {
		return nil, db.Error
	}

	return &user, nil
}
