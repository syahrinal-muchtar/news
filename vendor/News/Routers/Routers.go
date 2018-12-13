package Routers

import (
	"News/Controllers"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// r.Use(CORSMiddleware())

	// r.OPTIONS("/*path", CORSMiddleware())

	r.Static("/Public", "./Public") //Route Pictures

	r.POST("login", Controllers.LoginUser)

	// r.POST("upload", Controllers.UploadImgPost)

	v1 := r.Group("/api/v1/user")
	{
		v1.GET("/", auth, Controllers.ShowUsers)
		v1.GET("/:id", auth, Controllers.ShowUser)
		v1.POST("/", auth, Controllers.RegisterUser)
		v1.PUT("/:id", auth, Controllers.UpdateUser)
		v1.DELETE("/:id", auth, Controllers.DeleteUser)
	}
	v2 := r.Group("/api/v1/category")
	{
		v2.GET("/", auth, Controllers.ShowCategories)
		v2.GET("/:id", auth, Controllers.ShowCategory)
		v2.POST("/", auth, Controllers.AddCategory)
		v2.PUT("/:id", auth, Controllers.UpdateCategory)
		v2.DELETE("/:id", auth, Controllers.DeleteCategory)
	}
	v3 := r.Group("/api/v1/post")
	{
		v3.GET("/", auth, Controllers.ShowPosts)
		v3.GET("/:id", auth, Controllers.ShowPost)
		v3.POST("/", auth, Controllers.AddPost)
		v3.POST("upload", Controllers.UploadImgPost)
		v3.PUT("/:id", auth, Controllers.UpdatePost)
		v3.DELETE("/:id", auth, Controllers.DeletePost)
	}
	v4 := r.Group("/api/v1/tag")
	{
		v4.GET("/", auth, Controllers.ShowTags)
		v4.GET("/:id", auth, Controllers.ShowTag)
		v4.POST("/", auth, Controllers.AddTag)
		v4.PUT("/:id", auth, Controllers.UpdateTag)
		v4.DELETE("/:id", auth, Controllers.DeleteTag)
	}
	v5 := r.Group("/api/v1/post_tag")
	{
		v5.POST("/", auth, Controllers.AddPostTag)
		v5.GET("/:id", auth, Controllers.GetPost) //Show any post by tag
	}
	v6 := r.Group("/api/v1/news")
	{
		v6.GET("/", Controllers.ShowNews)
		v6.GET("/:id", Controllers.ShowNewsDetail)
	}
	v7 := r.Group("/api/v1/hotnews")
	{
		v7.GET("/", Controllers.ShowHotNews)
		v7.GET("/:id", Controllers.ShowHotNewsDetail)
	}

	return r
}

func preflight(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(http.StatusOK, struct{}{})
}

func auth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})

	// if token.Valid && err == nil {
	if token != nil && err == nil {
		fmt.Println("token verified")
	} else {
		result := gin.H{
			"message": "not authorized",
			"error":   err.Error(),
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}
}
