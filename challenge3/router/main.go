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
	doc, err := openapi3.NewLoader().LoadFromFile("../api/api.yaml")
	if (err != nil){
		fmt.Print(err.Error())
	}

	templateRouter, _ = legacy.NewRouter(doc)

	userRouter := router.Group("/user-management/user")
	initUserRouter(userRouter)

	roleRouter := router.Group("/user-management/role")
	initRoleRouter(roleRouter)

	grantRouter := router.Group("/user-management/grant")
	initGrantRouter(grantRouter)

	permissionRouter := router.Group("/user-management/permission")
	initPermissionRouter(permissionRouter)
}
