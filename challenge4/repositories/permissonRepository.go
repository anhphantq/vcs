package repositories

import (
	"challenge4/models"

	"gorm.io/gorm"
)

type PermissionRepository interface {
	SavePermission(permisson models.Permission) (models.Permission, error)
	FindPermission(permission models.Permission) ([]models.Permission, error)
	DeletePermission(models.Permission) error
	UpdatePermission(permission models.Permission) (models.Permission, error)
}

type IPermissionRepository struct {
	DB *gorm.DB
}

func (repo *IPermissionRepository) SavePermission(Permission models.Permission) (models.Permission, error) {
	result := repo.DB.Model(&models.Permission{}).Create(&Permission)

	if result.Error != nil {
		return models.Permission{}, result.Error
	}

	return Permission, nil
}
func (repo *IPermissionRepository) FindPermission(Permission models.Permission) ([]models.Permission, error) {
	var Permissions []models.Permission

	result := repo.DB.Model(&models.Permission{}).Where(&Permission).Scan(&Permissions)

	if result.Error != nil {
		return nil, result.Error
	}

	return Permissions, nil
}
func (repo *IPermissionRepository) DeletePermission(Permission models.Permission) error {
	result := repo.DB.Model(&models.Permission{}).Delete(&Permission)

	return result.Error
}
func (repo *IPermissionRepository) UpdatePermission(Permission models.Permission) (models.Permission, error) {
	result := repo.DB.Save(&Permission)

	if result.Error != nil {
		return models.Permission{}, result.Error
	}

	return Permission, nil
}
