package Models

import (
	"News/Config"

	_ "github.com/go-sql-driver/mysql"
)

func AddTag(b *Tag) (err error) {
	if err = Config.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func ShowTags(b *[]Tag) (err error) {
	if err = Config.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func ShowTag(b *Tag, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTag(b *Tag, id string) (err error) {
	Config.DB.Save(b)
	return nil
}

func DeleteTag(b *Tag, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(b)
	return nil
}
