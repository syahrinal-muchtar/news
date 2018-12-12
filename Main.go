package main

import (
	"News/Models"
	"encoding/json"
	"fmt"

	"News/Routers"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

var err error

type coba struct {
	Title  string
	Author string
}

func main22() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// val, err := client.Get("latestnews").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("latestnews", val)

	// err = client.Set("latestnews", "valueapa coba", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	val, err := client.LRange("latestNews", 0, 100).Result()
	if err != nil {
		panic(err)
	}

	b, _ := json.Marshal(val)
	// Convert bytes to string.
	s := string(b)
	fmt.Println(s)

	for i := range val {
		fmt.Println(val[i])
	}
	// mapD := map[string]int{"apple": 5, "lettuce": 7}
	// mapD = map[string]int{"apple": 5, "lettuce": 7}
	// x["key"] = append(x["key"], "value")
	// mapB, _ := json.Marshal(mapD)
	// fmt.Println(string(mapB))

	// x := make(map[string][]string)

	// x["key"] = append(x["key"], "value")
	// x["key"] = append(x["key"], "value1")

	// fmt.Println(x["key"][0])
	// fmt.Println(x["key"][1])

	// fmt.Println("latestnews", val[0])
	// fmt.Println(ini.Title)
}

func main() {

	DB, err := gorm.Open("mysql", "b181e8e3a141e0:f4928a04@tcp(us-cdbr-iron-east-01.cleardb.net)/heroku_ea827574aff1230?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("status: ", err)
	}
	defer DB.Close()

	//Migration User
	DB.AutoMigrate(&Models.User{})
	DB.Model(&Models.User{}).AddUniqueIndex("idx_email", "email")

	//Migration Category
	DB.AutoMigrate(&Models.Category{})
	DB.Model(&Models.Category{}).AddUniqueIndex("idx_name", "name")

	//Migration POST
	DB.AutoMigrate(&Models.Post{})
	DB.Model(&Models.Post{}).ModifyColumn("content", "text")
	DB.Model(&Models.Post{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	DB.Model(&Models.Post{}).AddForeignKey("category_id", "categories(id)", "RESTRICT", "RESTRICT")

	//Migration Comment
	DB.AutoMigrate(&Models.Comment{})
	DB.Model(&Models.Comment{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")

	//Migration Tag
	DB.AutoMigrate(&Models.Tag{})
	//Migration Post_Tag
	DB.AutoMigrate(&Models.Post_Tag{})
	DB.Model(&Models.Post_Tag{}).AddForeignKey("post_id", "posts(id)", "NO ACTION", "NO ACTION")
	DB.Model(&Models.Post_Tag{}).AddForeignKey("tag_id", "tags(id)", "NO ACTION", "NO ACTION")

	r := Routers.SetupRouter()
	// running
	r.Run()
}
