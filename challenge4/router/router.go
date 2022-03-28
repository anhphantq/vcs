package router

import (
	"challenge4/services"
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/routers"
	"github.com/getkin/kin-openapi/routers/legacy"
	"github.com/gin-gonic/gin"
)

var TemplateRouter routers.Router

func InitRouter(router *gin.Engine, userService services.UserService, roleService services.RoleService, permissionService services.PermissionService, grantService services.GrantService, postSerivce services.PostService) {
	doc, err := openapi3.NewLoader().LoadFromFile("/Users/anhphantq/Desktop/Go/vcs/challenge3/api/api.yaml")
	if err != nil {
		fmt.Print(err.Error())
	}

	TemplateRouter, err = legacy.NewRouter(doc)

	if err != nil {
		fmt.Print(err.Error())
	}

	userRouter := router.Group("/user-management")
	InitUserRouter(userRouter, userService)

	roleRouter := router.Group("/role-management")
	InitRoleRouter(roleRouter, userService, roleService)

	grantRouter := router.Group("/granting-management")
	InitGrantRouter(grantRouter, userService, grantService)

	permissionRouter := router.Group("/permission-management")
	InitPermissionRouter(permissionRouter, userService, permissionService)

	postRouter := router.Group("/post-management")
	InitPostRouter(postRouter, userService, postSerivce)
}
