package services

import (
	"challenge3/models"
)

type UserService interface {
	// User
	CheckEmailUsed(email string) (bool, error)
	SaveUser(user models.Account) (models.Account, error)
	CheckAuthentication(auth models.Authentication) (bool, error)
	CheckPermission(user_id uint, permissionName, scope string) (bool, error)
	CheckRole(user_id uint, role string) (bool, error)
	GetUserByID(id uint) (models.Account, error)
	GetUserByEmail(email string) (models.Account, error)
	GetAllUser() ([]models.Account, error)
	DeleteUserByID(id uint) error
	UpdateUser(user models.Account) (models.Account, error)
}

type RoleService interface {
	// Role
	GetAllRoles() ([]models.Role, error)
	GetRoleByID(id uint) (models.Account, error)
	SaveRole(role models.Role) (models.Role, error)
	DeleteRoleByID(id uint) error
}

type PermissionService interface {
	// Permisson
	GetAllPermisson() ([]models.Permission, error)
	GetPermissonByID(id uint) (models.Permission, error)
	SavePermission(permission models.Permission) (models.Permission, error)
	DeletePermissionByID(id uint) error
}

type GrantService interface {
	// Grant
	GetAllGrant() ([]models.RolePermission, error)
	DeleteGrant(models.RolePermission) error
	SaveGrant(models.RolePermission) (models.RolePermission, error)
}

type PostService interface {
}
