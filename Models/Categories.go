package Models

import (
	"News/Config"
)

func ShowCategories(b *[]Category) (err error) {
	if err = Config.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func ShowCategory(b *Category, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func AddCategory(b *Category) (err error) {
	if err = Config.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func UpdateCategory(b *Category, id string) (err error) {
	Config.DB.Save(b)
	return nil
}

func DeleteCategory(b *Category, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(b)
	return nil
}
