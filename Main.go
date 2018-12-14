package main

import (
	"News/Models"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"News/Config"
	// "News/Routers"
	"News/Controllers"
)

var err error

func main() {

	Config.DB, err = gorm.Open("mysql", "b181e8e3a141e0:f4928a04@tcp(us-cdbr-iron-east-01.cleardb.net:3306)/heroku_ea827574aff1230?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("status: ", err)
	}
	defer Config.DB.Close()

	//Migration User
	Config.DB.AutoMigrate(&Models.User{})
	Config.DB.Model(&Models.User{}).AddUniqueIndex("idx_email", "email")

	//Migration Category
	Config.DB.AutoMigrate(&Models.Category{})
	Config.DB.Model(&Models.Category{}).AddUniqueIndex("idx_name", "name")

	//Migration POST
	Config.DB.AutoMigrate(&Models.Post{})
	Config.DB.Model(&Models.Post{}).ModifyColumn("content", "text")
	Config.DB.Model(&Models.Post{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	Config.DB.Model(&Models.Post{}).AddForeignKey("category_id", "categories(id)", "RESTRICT", "RESTRICT")

	//Migration Comment
	Config.DB.AutoMigrate(&Models.Comment{})
	Config.DB.Model(&Models.Comment{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")

	//Migration Tag
	Config.DB.AutoMigrate(&Models.Tag{})
	//Migration Post_Tag
	Config.DB.AutoMigrate(&Models.Post_Tag{})
	Config.DB.Model(&Models.Post_Tag{}).AddForeignKey("post_id", "posts(id)", "NO ACTION", "NO ACTION")
	Config.DB.Model(&Models.Post_Tag{}).AddForeignKey("tag_id", "tags(id)", "NO ACTION", "NO ACTION")

	// r := Routers.SetupRouter()

	r := gin.Default()
	r.Use(cors.Default())

	// r.Use(cors.New(cors.Config{
	// 	AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT"},
	// 	AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	AllowAllOrigins:  false,
	// 	AllowOriginFunc:  func(origin string) bool { return true },
	// 	MaxAge:           86400,
	// }))

	r.Static("/Public", "./Public") //Route Pictures

	r.POST("login", Controllers.LoginUser)

	// r.POST("upload", Controllers.UploadImgPost)
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

	// running
	r.Run()
}
