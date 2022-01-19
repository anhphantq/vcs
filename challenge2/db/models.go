package db

type Account struct {
	User_id  uint   `json:"user_id" gorm:"primaryKey"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"contains=@"`
	Role_id  uint   `json:"role_id,string" validate:"number"`
	Password string `json:"password" validate:"min=8"`
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
	Role_id       uint `json:"role_id,string"`
	Permission_id uint `json:"permission_id,string"`
}

type Authentication struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
