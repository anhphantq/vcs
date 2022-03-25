package repositories

import "challenge3/models"

type GrantRepository interface {
	SaveGrant(models.RolePermission) (models.RolePermission, error)
	FindGrant(models.RolePermission) ([]models.RolePermission, error)
	DeleteGrant(models.RolePermission) error
}

type IGrantRepository struct {
}
