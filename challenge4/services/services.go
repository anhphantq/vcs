package services

import (
	"challenge4/models"
)

type UserService interface {
	// User
	CheckEmailUsed(email string) (bool, error)
	SaveUser(user models.Account) (models.Account, error)
	//CheckAuthentication(auth models.Authentication) (bool, error)
	CheckPermission(user_id uint, permissionName, scope string) (bool, error)
	CheckRole(user_id uint, role string) (bool, error)
	GetUserByID(id uint) (models.Account, error)
	GetUserByEmail(email string) (models.Account, error)
	GetAllUser() ([]models.Account, error)
	DeleteUserByID(id uint) error
	UpdateUserByID(user models.Account) (models.Account, error)
}

type RoleService interface {
	// Role
	GetAllRole() ([]models.Role, error)
	GetRoleByID(id uint) (models.Role, error)
	InsertRole(role models.Role) (models.Role, error)
	DeleteRoleByID(id uint) error
}

type PermissionService interface {
	// Permisson
	GetAllPermission() ([]models.Permission, error)
	GetPermissionByID(id uint) (models.Permission, error)
	InsertPermission(permission models.Permission) (models.Permission, error)
	DeletePermissionByID(id uint) error
}

type GrantService interface {
	// Grant
	GetAllGrant() ([]models.Rolepermission, error)
	DeleteGrantSrv(models.Rolepermission) error
	InsertGrant(models.Rolepermission) (models.Rolepermission, error)
}

type PostService interface {
	GetPostByUserID(user_id uint, page int) ([]models.Post, error)
	GetPostByID(post_id uint) (models.Post, error)
	InsertPost(models.Post) (models.Post, error)
	UpdatePostSrv(post models.Post) (models.Post, error)
	DeletePostSrv(post_id uint) error
	//admin
	GetAllPosts(page int) ([]models.Post, error)
}
