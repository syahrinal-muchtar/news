package main

import (
	"News/Models"
	"fmt"

	"github.com/jinzhu/gorm"

	"News/Config"
	"News/Routers"
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
