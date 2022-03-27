package main

import (
	"challenge4/db"
	"challenge4/repositories"
	"challenge4/router"
	"challenge4/services"

	"github.com/gin-gonic/gin"
)

func main() {
	mainRouter := gin.Default()

	userRepo := &repositories.IUserRepository{DB: db.GetDB()}
	roleRepo := &repositories.IRoleRepository{DB: db.GetDB()}
	permissionRepo := &repositories.IPermissionRepository{DB: db.GetDB()}
	grantRepo := &repositories.IGrantRepository{DB: db.GetDB()}
	postRepo := &repositories.IPostRepository{DB: db.GetDB()}

	userService := &services.IUserSerivce{UserRepository: userRepo, RoleRepository: roleRepo, PermissionRepository: permissionRepo, GrantRepository: grantRepo}
	roleService := &services.IRoleService{RoleRepository: roleRepo}
	permissionService := &services.IPermissionService{PermissionRepository: permissionRepo}
	grantService := &services.IGrantService{GrantRepository: grantRepo}
	postService := &services.IPostService{PostRepository: postRepo}

	router.InitRouter(mainRouter, userService, roleService, permissionService, grantService, postService)
	mainRouter.Run(":8080")
}
