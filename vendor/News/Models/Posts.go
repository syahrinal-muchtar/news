package Models

import (
	"News/Config"

	_ "github.com/go-sql-driver/mysql"
)

func AddPost(b *Post) (err error) {
	if err = Config.DB.Create(b).Error; err != nil {
		return err
	}
	// Config.DB.Preload("User").Find(b)
	// Config.DB.Preload("Category").Find(b)
	return nil
}

func ShowPosts(b *[]Post) (err error) {
	if err = Config.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func ShowPost(b *Post, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func ShowPostByCategory(b *Post, catId string) (err error) {
	if err := Config.DB.Where("category_id = ?", catId).First(b).Error; err != nil {
		return err
	}
	return nil
}

func UpdatePost(b *Post, id string) (err error) {
	Config.DB.Save(b)
	return nil
}

func DeletePost(b *Post, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(b)
	return nil
}
