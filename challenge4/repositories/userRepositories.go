package repositories

import "challenge3/models"

type UserRepository interface {
	SaveUser(user models.Account) (models.Account, error)
	FindUser(user models.Account) ([]models.Account, error)
	DeleteUserByID(id uint) error
	UpdateUserByID(id uint, user models.Account) (models.Account, error)
}

type IUserRepository struct {
}
