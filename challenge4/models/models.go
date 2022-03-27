package models

type Account struct {
	User_id  uint   `json:"user_id" gorm:"primaryKey"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role_id  uint   `json:"role_id"`
	Password string `json:"password"`
}

type Role struct {
	Role_id uint   `json:"role_id,string" gorm:"primaryKey"`
	Name    string `json:"name" binding:"required"`
}

type Permission struct {
	Permission_id uint   `json:"permission_id,string" gorm:"primaryKey"`
	Name          string `json:"name"`
	Scope         string `json:"scope"`
}

type Rolepermission struct {
	Role_id       uint `json:"role_id" form:"role_id" gorm:"primaryKey"`
	Permission_id uint `json:"permission_id" form:"permission_id" gorm:"primaryKey"`
}

type Authentication struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Post struct {
	Post_id    uint   `json:"post_id" gorm:"primaryKey"`
	Content    string `json:"content"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	User_id    uint   `json:"user_id"`
}
