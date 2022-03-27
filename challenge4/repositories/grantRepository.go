package repositories

import (
	"challenge4/models"

	"gorm.io/gorm"
)

type GrantRepository interface {
	SaveGrant(models.Rolepermission) (models.Rolepermission, error)
	FindGrant(models.Rolepermission) ([]models.Rolepermission, error)
	DeleteGrant(models.Rolepermission) error
}

type IGrantRepository struct {
	DB *gorm.DB
}

func (repo *IGrantRepository) SaveGrant(Rolepermission models.Rolepermission) (models.Rolepermission, error) {
	result := repo.DB.Model(&models.Rolepermission{}).Create(&Rolepermission)

	if result.Error != nil {
		return models.Rolepermission{}, result.Error
	}

	return Rolepermission, nil
}
func (repo *IGrantRepository) FindGrant(Rolepermission models.Rolepermission) ([]models.Rolepermission, error) {
	var Rolepermissions []models.Rolepermission

	result := repo.DB.Model(&models.Rolepermission{}).Where(&Rolepermission).Scan(&Rolepermissions)

	if result.Error != nil {
		return nil, result.Error
	}

	return Rolepermissions, nil
}
func (repo *IGrantRepository) DeleteGrant(Rolepermission models.Rolepermission) error {
	result := repo.DB.Model(&models.Rolepermission{}).Delete(&Rolepermission)

	return result.Error
}
