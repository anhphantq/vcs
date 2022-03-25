package repositories

import "challenge3/models"

type PermissonRepository interface {
	SavePermission(permisson models.Permission) (models.Permission, error)
	FindPermisson(permission models.Permission) ([]models.Permission, error)
	DeletePermissonByID(id uint) error
	UpdatePermissonByID(id uint, permission models.Permission) (models.Permission, error)
}

type IPermissonRepository struct {
}
