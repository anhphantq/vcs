package repositories

import (
	"challenge4/models"

	"gorm.io/gorm"
)

type RoleRepository interface {
	SaveRole(role models.Role) (models.Role, error)
	FindRole(role models.Role) ([]models.Role, error)
	DeleteRole(role models.Role) error
	UpdateRole(role models.Role) (models.Role, error)
}

type IRoleRepository struct {
	DB *gorm.DB
}

func (repo *IRoleRepository) SaveRole(role models.Role) (models.Role, error) {
	result := repo.DB.Model(&models.Role{}).Create(&role)

	if result.Error != nil {
		return models.Role{}, result.Error
	}

	return role, nil
}
func (repo *IRoleRepository) FindRole(role models.Role) ([]models.Role, error) {
	var roles []models.Role

	result := repo.DB.Model(&models.Role{}).Where(&role).Scan(&roles)

	if result.Error != nil {
		return nil, result.Error
	}

	return roles, nil
}
func (repo *IRoleRepository) DeleteRole(role models.Role) error {
	result := repo.DB.Model(&models.Role{}).Delete(&role)

	return result.Error
}
func (repo *IRoleRepository) UpdateRole(role models.Role) (models.Role, error) {
	result := repo.DB.Save(&role)

	if result.Error != nil {
		return models.Role{}, result.Error
	}

	return role, nil
}
