package middleware

import (
	"net/http"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers"
	"github.com/gin-gonic/gin"
)

func ValidationMiddleware(templateRouter routers.Router) gin.HandlerFunc{
	return func(c *gin.Context) {
		route, pathParams, err := templateRouter.FindRoute(c.Request)

		if (err != nil){
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
			return
		}

		requestValidationInput := &openapi3filter.RequestValidationInput{
			Request:    c.Request,
			PathParams: pathParams,
			Route:      route,
			Options: &openapi3filter.Options{
				AuthenticationFunc: openapi3filter.NoopAuthenticationFunc,
			},
		}

		if err := openapi3filter.ValidateRequest(c.Request.Context(), requestValidationInput); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.Next()
	}
}