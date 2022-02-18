package db

type Account struct {
	User_id  uint   `json:"user_id" gorm:"primaryKey"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role_id  uint   `json:"role_id"`
	Password string `json:"password"`
}

type Role struct {
	Role_id uint   `json:"role_id,string"`
	Name    string `json:"name" binding:"required"`
}

type Permission struct {
	Permission_id uint   `json:"permission_id,string"`
	Name          string `json:"name"`
	Scope         string `json:"scope"`
}

type RolePermission struct {
	Role_id       uint `json:"role_id" form:"role_id"`
	Permission_id uint `json:"permission_id" form:"permission_id"`
}

type Authentication struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
