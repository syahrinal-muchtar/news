package Controllers

import (
	"News/ApiHelpers"
	"News/Config"
	"News/Models"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func ShowNews(c *gin.Context) {
	page := c.Query("page")
	var news []Models.News
	err := Models.ShowNews(&news)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, news)
	} else {
		page, _ := strconv.Atoi(c.DefaultQuery("page", "2"))
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))

		paginator := pagination.Pagging(&pagination.Param{
			DB:      Config.DB,
			Page:    page,
			Limit:   limit,
			OrderBy: []string{"id desc"},
			ShowSQL: true,
		}, &news)

		ApiHelpers.RespondJSON(c, 200, paginator)
	}
}

func ShowNewsDetail(c *gin.Context) {
	id := c.Params.ByName("id")
	var news Models.News
	err := Models.ShowNewsDetail(&news, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, news)
	} else {
		ApiHelpers.RespondJSON(c, 200, news)
	}
}

func ShowHotNews(c *gin.Context) {
	var resolvedURL = os.Getenv("REDIS_URL")
	var password = ""
	if !strings.Contains(resolvedURL, "localhost") {
		parsedURL, _ := url.Parse(resolvedURL)
		password, _ = parsedURL.User.Password()
		resolvedURL = parsedURL.Host
	}
	fmt.Printf("connecting to %s", resolvedURL)
	client := redis.NewClient(&redis.Options{
		Addr:     resolvedURL,
		Password: password,
		DB:       0,
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

	news := []Models.HotNews{}
	for i := range title {
		n := Models.HotNews{ID: i + 1, Title: title[i], Date: date[i], Content: content[i], Image: image[i], Author: author[i]}
		news = append(news, n)
	}

	c.JSON(200, news)
}

func ShowHotNewsDetail(c *gin.Context) {
	// news := []Models.HotNews{}
	// err := Models.ShowHotNews(&news, id)

	// ApiHelpers.RespondJSON(c, 200, news)
	// ids := c.Query("id")
	ids := c.PostForm("id")
	id, err := strconv.Atoi(ids)

	if err != nil {
		// handle error
	}
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
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

	news := []Models.HotNews{}
	for i := range title {
		n := Models.HotNews{ID: i + 1, Title: title[i], Date: date[i], Content: content[i], Image: image[i], Author: author[i]}
		news = append(news, n)
	}

	for i := range news {
		if news[i].ID == id {
			c.JSON(200, news[i])
			break
		}
	}
}
