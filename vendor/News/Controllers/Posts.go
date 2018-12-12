package Controllers

import (
	"News/ApiHelpers"
	"News/Config"
	"News/Models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func ShowPosts(c *gin.Context) {
	var post []Models.Post
	err := Models.ShowPosts(&post)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, post)
	} else {
		page, _ := strconv.Atoi(c.DefaultQuery("page", "3"))
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "3"))

		paginator := pagination.Pagging(&pagination.Param{
			DB:      Config.DB,
			Page:    page,
			Limit:   limit,
			OrderBy: []string{"updated_at desc"},
			ShowSQL: true,
		}, &post)
		// c.JSON(200, paginator)
		ApiHelpers.RespondJSON(c, 200, paginator)
	}
}

func ShowPost(c *gin.Context) {
	// id := c.Params.ByName("id")
	// var post Models.Post
	// err := Models.ShowPost(&post, id)
	// if err != nil {
	// 	ApiHelpers.RespondJSON(c, 404, post)
	// } else {
	// 	ApiHelpers.RespondJSON(c, 200, post)
	// }

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	title, err := client.LRange("title", 0, 100).Result()
	if err != nil {
		panic(err)
	}
	date, err := client.LRange("date", 0, 100).Result()
	if err != nil {
		panic(err)
	}
	author, err := client.LRange("author", 0, 100).Result()
	if err != nil {
		panic(err)
	}
	image, err := client.LRange("image", 0, 100).Result()
	if err != nil {
		panic(err)
	}
	content, err := client.LRange("content", 0, 100).Result()
	if err != nil {
		panic(err)
	}

	// b, _ := json.Marshal(val)
	// // Convert bytes to string.
	// s := string(b)

	for i := range title {
		fmt.Println(title[i])
		fmt.Println(date[i])
		fmt.Println(author[i])
		fmt.Println(image[i])
		fmt.Println(content[i])
	}

	// json_string := `{"firstname": "Rocky","lastname": "Sting","city": "London"}`

	// b := []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)

	// in := []byte(val[0])
	// var raw map[string]interface{}
	// json.Unmarshal(in, &raw)
	// raw["count"] = 1
	// out, _ := json.Marshal(raw)
	// println(string(out))
	// out := json.Marshal("/"")
	// arr := []string{"test1", "test3", "test2"}

	// var news []Models.News
	// news := new([]Models.News)

	// emp6 := Models.News{"Sam", "Anderson", "test", "test"}

	// news.Title = "test"
	// news[0].Content = "test"
	// news[0].Image = "tesf  fgrrt"
	// news[0].Author = "test"

	news := []Models.News{}
	// for i := 0; i < 47; i++ {
	// 	n := Models.News{ID: i, Title: "test", Content: "test", Image: "test", Author: "test"}
	// 	news = append(news, n)
	// }

	for i := range title {
		n := Models.News{ID: i + 1, Title: title[i], Content: content[i], Image: image[i], Author: author[i]}
		news = append(news, n)
	}

	// news.Date = "test"

	c.JSON(200, news)

	// ApiHelpers.RespondJSON(c, 200, out)
	// c.JSON(200, gin.H{
	// 	"status":  "posted",
	// 	"message": "message",
	// 	"nick":    "nick",
	// })
}

func ShowPostByCategory(c *gin.Context) {
	id := c.Params.ByName("id")
	var post Models.Post
	err := Models.ShowPostByCategory(&post, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, post)
	} else {
		ApiHelpers.RespondJSON(c, 200, post)
	}
}

func AddPost(c *gin.Context) {
	// file, err := c.FormFile("picture")
	// if err != nil {
	// 	c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
	// 	return
	// }

	// path := fmt.Sprintf("Public/Pictures/Posts/%s", file.Filename)
	// if err := c.SaveUploadedFile(file, path); err != nil {
	// 	c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
	// 	return
	// }

	// post.Title = c.PostForm("title")
	// post.Content = c.PostForm("content")

	// post.UserID = StringtoUint(c.PostForm("user_id"))

	// post.CategoryID = StringtoUint(c.PostForm("category_id"))
	// post.Picture = fmt.Sprintf("localhost:8080/%s", path)

	var post Models.Post

	c.BindJSON(&post)

	err := Models.AddPost(&post)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, post)
	} else {
		ApiHelpers.RespondJSON(c, 200, post)
	}
}

func StringtoUint(id string) uint {
	u64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	userID := uint(u64)
	return userID
}

func UploadImgPost(c *gin.Context) {
	file, err := c.FormFile("picture")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	path := fmt.Sprintf("Public/Pictures/Posts/%s", file.Filename)
	if err := c.SaveUploadedFile(file, path); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("localhost:8080/%s", path))
}

func UpdatePost(c *gin.Context) {
	var post Models.Post
	id := c.Params.ByName("id")
	err := Models.ShowPost(&post, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, post)
	}
	c.BindJSON(&post)
	err = Models.UpdatePost(&post, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, post)
	} else {
		ApiHelpers.RespondJSON(c, 200, post)
	}
}

func DeletePost(c *gin.Context) {
	var post Models.Post
	id := c.Params.ByName("id")
	err := Models.DeletePost(&post, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, post)
	} else {
		ApiHelpers.RespondJSON(c, 200, post)
	}
}
