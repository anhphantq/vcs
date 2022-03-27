package router

import (
	"challenge4/middleware"
	"challenge4/models"
	"challenge4/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var url = "http://localhost:8080"
var postService services.PostService

func hdGetPost(c *gin.Context) {
	tmp, _ := c.Get("user")
	user, _ := tmp.(models.Account)

	page := c.Query("page")

	if page == "" {
		c.Redirect(http.StatusPermanentRedirect, url+"/post-management/post?page=0")
		return
	}

	pageID, err := strconv.Atoi(page)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong with query params"})
		return
	}

	posts, err := postService.GetPostByUserID(user.User_id, pageID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Wrong page number or something went wrong in the server"})
		return
	}

	c.JSON(http.StatusAccepted, posts)
}

func hdCreatePost(c *gin.Context) {
	var post models.Post

	err := c.ShouldBind(&post)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"messsage": "Something went wrong in the server"})
		return
	}

	tmp, _ := c.Get("user")
	user, _ := tmp.(models.Account)

	post.User_id = user.User_id

	post, err = postService.InsertPost(post)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	c.JSON(http.StatusAccepted, post)
}

func hdUpdatePostByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	tmp, _ := c.Get("user")
	user, _ := tmp.(models.Account)

	var postTMP models.Post

	err = c.ShouldBind(&postTMP)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	var post models.Post

	post, err = postService.GetPostByID(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server or post does not exists"})
		return
	}

	if post.User_id != user.User_id {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Post does not belong to user"})
		return
	}

	post.Content = postTMP.Content

	post, err = postService.UpdatePostSrv(post)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server or wrong post id"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Post updated"})
}

func hdGetPostByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	tmp, _ := c.Get("user")
	user, _ := tmp.(models.Account)

	var post models.Post

	post, err = postService.GetPostByID(uint(id))

	if err != nil || post.User_id != user.User_id {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server or wrong post's ID"})
		return
	}

	c.JSON(http.StatusAccepted, post)
}

func hdDeletePostByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	tmp, _ := c.Get("user")
	user, _ := tmp.(models.Account)

	post, err := postService.GetPostByID(uint(id))

	if err != nil || post.User_id != user.User_id {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server or wrong post's ID"})
		return
	}

	err = postService.DeletePostSrv(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server or wrong post id"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Post deleted"})
}

func hdGetPosts(c *gin.Context) {
	var posts []models.Post

	page := c.Query("page")

	if page == "" {
		c.Redirect(http.StatusPermanentRedirect, url+"/post-management/posts?page=0")
		return
	}

	pageID, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	posts, err = postService.GetAllPosts(pageID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Wrong page number or something went wrong in the server"})
		return
	}

	c.JSON(http.StatusAccepted, posts)
}

func hdGetPostsByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	var post models.Post

	post, err = postService.GetPostByID(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server or wrong post's ID"})
		return
	}

	c.JSON(http.StatusAccepted, post)
}

func hdUpdatePostsByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	var postTMP models.Post

	err = c.ShouldBind(&postTMP)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	var post models.Post

	post, err = postService.GetPostByID(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server or post does not exists"})
		return
	}

	post.Content = postTMP.Content

	post, err = postService.UpdatePostSrv(post)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server or wrong post id"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Post updated"})
}

func hdDeletePostsByID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	err = postService.DeletePostSrv(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server or wrong post id"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Post deleted"})
}

func InitPostRouter(router *gin.RouterGroup, userService services.UserService, postservice services.PostService) {
	postService = postservice
	router.GET("/post", middleware.ValidationMiddleware(templateRouter), middleware.AuthMiddleware(userService), hdGetPost)
	router.POST("/post", middleware.ValidationMiddleware(templateRouter), middleware.AuthMiddleware(userService), hdCreatePost)

	router.GET("/post/:id", middleware.ValidationMiddleware(templateRouter), middleware.AuthMiddleware(userService), hdGetPostByID)
	router.PUT("/post/:id", middleware.ValidationMiddleware(templateRouter), middleware.AuthMiddleware(userService), hdUpdatePostByID)
	router.DELETE("/post/:id", middleware.ValidationMiddleware(templateRouter), middleware.AuthMiddleware(userService), hdDeletePostByID)

	router.GET("/posts", middleware.ValidationMiddleware(templateRouter), middleware.AuthMiddleware(userService), middleware.RoleValidationMiddleware(userService, "admin"), hdGetPosts)
	router.GET("/posts/:id", middleware.ValidationMiddleware(templateRouter), middleware.AuthMiddleware(userService), middleware.RoleValidationMiddleware(userService, "admin"), hdGetPostsByID)
	router.PUT("/posts/:id", middleware.ValidationMiddleware(templateRouter), middleware.AuthMiddleware(userService), middleware.RoleValidationMiddleware(userService, "admin"), hdUpdatePostsByID)
	router.DELETE("/posts/:id", middleware.ValidationMiddleware(templateRouter), middleware.AuthMiddleware(userService), middleware.RoleValidationMiddleware(userService, "admin"), hdDeletePostsByID)
}
