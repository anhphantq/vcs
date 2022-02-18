package router

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/routers"
	"github.com/getkin/kin-openapi/routers/legacy"
	"github.com/gin-gonic/gin"
)

var templateRouter routers.Router

func InitRouter(router *gin.Engine) {
	doc, err := openapi3.NewLoader().LoadFromFile("/Users/anhphantq/Desktop/Go/vcs/challenge3/api/api.yaml")
	if (err != nil){
		fmt.Print(err.Error())
	}

	templateRouter, err = legacy.NewRouter(doc)

	if (err != nil){
		fmt.Print(err.Error())
	}

	userRouter := router.Group("/user-management")
	initUserRouter(userRouter)

	roleRouter := router.Group("/role-management")
	initRoleRouter(roleRouter)

	grantRouter := router.Group("/granting-management")
	initGrantRouter(grantRouter)

	permissionRouter := router.Group("/permission-management")
	initPermissionRouter(permissionRouter)
}
