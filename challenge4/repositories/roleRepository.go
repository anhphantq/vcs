package repositories

import "challenge3/models"

type RoleRepository interface {
	SaveRole(role models.Role) (models.Role, error)
	FindRole(role models.Role) ([]models.Role, error)
	DeleteRoleByID(id uint) error
	UpdateRoleByID(id uint, role models.Role) (models.Role, error)
}

type IRoleRepository struct {
}
