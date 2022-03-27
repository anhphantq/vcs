package repositories

import (
	"challenge4/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user models.Account) (models.Account, error)
	FindUser(user models.Account) ([]models.Account, error)
	DeleteUser(user models.Account) error
	UpdateUser(user models.Account) (models.Account, error)
}

type IUserRepository struct {
	DB *gorm.DB
}

func (repo *IUserRepository) InsertUser(user models.Account) (models.Account, error) {
	result := repo.DB.Model(&models.Account{}).Create(&user)

	if result.Error != nil {
		return models.Account{}, result.Error
	}

	return user, nil
}

func (repo *IUserRepository) FindUser(user models.Account) ([]models.Account, error) {
	var users []models.Account

	result := repo.DB.Model(&models.Account{}).Where(&user).Scan(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (repo *IUserRepository) DeleteUser(user models.Account) error {
	result := repo.DB.Model(&models.Account{}).Delete(&user)

	return result.Error
}

func (repo *IUserRepository) UpdateUser(user models.Account) (models.Account, error) {
	result := repo.DB.Save(&user)

	if result.Error != nil {
		return models.Account{}, result.Error
	}

	return user, nil
}
