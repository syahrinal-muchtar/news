package Models

import (
	"News/Config"

	_ "github.com/go-sql-driver/mysql"
)

func ShowNews(b *[]News) (err error) {
	if err = Config.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func ShowNewsDetail(b *News, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

// func ShowHotNews(b []HotNews) {
// 	client := redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6379",
// 		Password: "",
// 		DB:       0,
// 	})

// 	title, err := client.LRange("title", 0, 100).Result()
// 	if err != nil {
// 		panic(err)
// 	}
// 	date, err := client.LRange("date", 0, 100).Result()
// 	if err != nil {
// 		panic(err)
// 	}
// 	author, err := client.LRange("author", 0, 100).Result()
// 	if err != nil {
// 		panic(err)
// 	}
// 	image, err := client.LRange("image", 0, 100).Result()
// 	if err != nil {
// 		panic(err)
// 	}
// 	content, err := client.LRange("content", 0, 100).Result()
// 	if err != nil {
// 		panic(err)
// 	}

// 	news := []HotNews{}
// 	for i := range title {
// 		n := HotNews{ID: i + 1, Title: title[i], Date: date[i], Content: content[i], Image: image[i], Author: author[i]}
// 		&news = append(news, n)
// 	}
// }

func UpdateNews(b *News, id string) (err error) {
	Config.DB.Save(b)
	return nil
}

func DeleteNews(b *News, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(b)
	return nil
}
