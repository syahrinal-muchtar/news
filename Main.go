package main

import (
	"News/Models"
	"News/Routers"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"News/Config"
	// "News/Routers"
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

	r := Routers.SetupRouter()

	// running
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "https://news148.herokuapp.com")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
