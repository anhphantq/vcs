package db

type Account struct {
	User_id  uint   `json:"user_id"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Role_id  uint   `json:"role_id,string"`
	Password string `json:"password" binding:"required"`
}

type Role struct {
	Role_id uint   `json:"role_id"`
	Name    string `json:"name"`
}

type Permission struct {
	Permission_id uint   `json:"permission_id"`
	Name          string `json:"name"`
	Scope         string `json:"scope"`
}

type RolePermission struct {
	Role_id       uint `json:"role_id"`
	Permission_id uint `json:"permission_id"`
}

type Authentication struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
