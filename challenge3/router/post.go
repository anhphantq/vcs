package router

import (
	"challenge3/db"
	"challenge3/middleware"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var url = "http://localhost:8080"

func hdGetPost(c *gin.Context){
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	tmp, _ := c.Get("user")
	user, _ := tmp.(db.Account)

	var posts []db.Post

	result := connection.Raw("select * from posts where user_id = ?", user.User_id).Scan(&posts)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server or no posts founded"})
		return
	}

	page := c.Query("page")

	if (len(posts) > 10 && page == ""){
		c.Redirect(http.StatusPermanentRedirect, url + "/post-management/post?page=0")
		return
	}

	var numP int

	if (len(posts)%10 == 0){
		numP = len(posts)/10
	} else{
		numP = len(posts)/10 + 1
	}

	if (page != ""){
		page, err := strconv.Atoi(page)
		if (err != nil){
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
			return
		}
		
		if (page + 1 > numP){
			c.JSON(http.StatusBadRequest, gin.H{"message": "Page is not defined"})
			return
		}
			
		if (page + 1 == numP){
			posts = posts[10*page:]
		} else {
			posts = posts[10*page : 10*(page+1) - 1]
		}
	}

	c.JSON(http.StatusAccepted, posts)
}

func hdCreatePost(c *gin.Context){
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	var post db.Post

	err := c.ShouldBind(&post)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"messsage": "Something went wrong in the server"})
		return
	}

	tmp, _ := c.Get("user")
	user, _ := tmp.(db.Account)

	result := connection.Raw("insert into posts values(default,?,default,default,?) returning *", post.Content, user.User_id).Scan(&post)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	c.JSON(http.StatusAccepted, post)
}

func hdUpdatePostByID(c *gin.Context){
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	tmp, _ := c.Get("user")
	user, _ := tmp.(db.Account)

	var post db.Post

	err = c.ShouldBind(&post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	result := connection.Exec("update posts set content = ? where user_id = ? and post_id = ?", post.Content, user.User_id, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	if result.RowsAffected < 1{
		c.JSON(http.StatusBadRequest, gin.H{"message": "Wrong post's ID"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Post updated"})
}

func hdGetPostByID(c *gin.Context){
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	tmp, _ := c.Get("user")
	user, _ := tmp.(db.Account)

	var post db.Post

	result := connection.Raw("select * from posts where user_id = ? and post_id = ?", user.User_id, id).Scan(&post)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server or wrong post's ID"})
		return
	}

	c.JSON(http.StatusAccepted, post)
}

func hdDeletePostByID(c *gin.Context){
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	tmp, _ := c.Get("user")
	user, _ := tmp.(db.Account)

	result := connection.Exec("delete from posts where user_id = ? and post_id = ?", user.User_id, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	if result.RowsAffected < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Wrong post's ID"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Post deleted"})
}

func hdGetPosts(c *gin.Context){
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	var posts []db.Post

	result := connection.Raw("select * from posts").Scan(&posts)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	page := c.Query("page")

	if (len(posts) > 10 && page == ""){
		c.Redirect(http.StatusPermanentRedirect, url + "/post-management/posts?page=0")
		return
	}

	var numP int

	if (len(posts)%10 == 0){
		numP = len(posts)/10
	} else{
		numP = len(posts)/10 + 1
	}

	if (page != ""){
		page, err := strconv.Atoi(page)
		if (err != nil){
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
			return
		}
		
		if (page + 1 > numP){
			c.JSON(http.StatusBadRequest, gin.H{"message": "Page is not defined"})
			return
		}
			
		if (page + 1 == numP){
			posts = posts[10*page:]
		} else {
			posts = posts[10*page : 10*(page+1) - 1]
		}
	}

	c.JSON(http.StatusAccepted, posts)
}

func hdGetPostsByUserID(c *gin.Context){
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	var posts []db.Post

	result := connection.Raw("select * from posts where user_id = ?", id).Scan(&posts)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server or no posts founded"})
		return
	}

	page := c.Query("page")

	if (len(posts) > 10 && page == ""){
		c.Redirect(http.StatusPermanentRedirect, url + "/post-management/posts/" + c.Param("id") + "?page=0")
		return
	}

	var numP int

	if (len(posts)%10 == 0){
		numP = len(posts)/10
	} else{
		numP = len(posts)/10 + 1
	}

	if (page != ""){
		page, err := strconv.Atoi(page)
		if (err != nil){
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
			return
		}
		
		if (page + 1 > numP){
			c.JSON(http.StatusBadRequest, gin.H{"message": "Page is not defined"})
			return
		}
			
		if (page + 1 == numP){
			posts = posts[10*page:]
		} else {
			posts = posts[10*page : 10*(page+1) - 1]
		}
	}

	c.JSON(http.StatusAccepted, posts)
}

func hdUpdatePostsByID(c *gin.Context){
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	var post db.Post

	err = c.ShouldBind(&post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	result := connection.Exec("update posts set content = ? where post_id = ?", post.Content, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	if result.RowsAffected < 1{
		c.JSON(http.StatusBadRequest, gin.H{"message": "Wrong post's ID"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Post updated"})
}

func hdDeletePostsByID(c *gin.Context){
	connection := db.GetDatabase()
	defer db.Closedatabase(connection)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	result := connection.Exec("delete from posts where post_id = ?", id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong in the server"})
		return
	}

	if result.RowsAffected < 1{
		c.JSON(http.StatusBadRequest, gin.H{"message": "Wrong post's ID"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Post deleted"})
}

func InitPostRouter(router *gin.RouterGroup){
	router.GET("/post", middleware.ValidationMiddleware(templateRouter), middleware.AuthMiddleware(), hdGetPost)
	router.POST("/post", middleware.ValidationMiddleware(templateRouter), middleware.AuthMiddleware(), hdCreatePost)

	router.GET("/post/:id", middleware.ValidationMiddleware(templateRouter), middleware.AuthMiddleware(), hdGetPostByID)
	router.PUT("/post/:id", middleware.ValidationMiddleware(templateRouter), middleware.AuthMiddleware(), hdUpdatePostByID)
	router.DELETE("/post/:id", middleware.ValidationMiddleware(templateRouter), middleware.AuthMiddleware(), hdDeletePostByID)

	router.GET("/posts", middleware.ValidationMiddleware(templateRouter), middleware.AuthMiddleware(), middleware.AuthAdminMiddleware(), hdGetPosts)
	router.GET("/posts/:id", middleware.ValidationMiddleware(templateRouter), middleware.AuthMiddleware(), middleware.AuthAdminMiddleware(), hdGetPostsByUserID)
	router.PUT("/posts/:id", middleware.ValidationMiddleware(templateRouter), middleware.AuthMiddleware(), middleware.AuthAdminMiddleware(),hdUpdatePostsByID)
	router.DELETE("/posts/:id", middleware.ValidationMiddleware(templateRouter), middleware.AuthMiddleware(), middleware.AuthAdminMiddleware(), hdDeletePostsByID)
}