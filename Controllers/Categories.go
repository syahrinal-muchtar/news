package Controllers

import (
	"News/ApiHelpers"
	"News/Models"

	"github.com/gin-gonic/gin"
)

func ShowCategories(c *gin.Context) {
	var categories []Models.Category
	err := Models.ShowCategories(&categories)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, categories)
	} else {
		ApiHelpers.RespondJSON(c, 200, categories)
	}
}

func ShowCategory(c *gin.Context) {
	id := c.Params.ByName("id")
	var category Models.Category
	err := Models.ShowCategory(&category, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, category)
	} else {
		ApiHelpers.RespondJSON(c, 200, category)
	}
}

func AddCategory(c *gin.Context) {
	var category Models.Category
	c.BindJSON(&category)
	err := Models.AddCategory(&category)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, category)
	} else {
		ApiHelpers.RespondJSON(c, 200, category)
	}
}

func UpdateCategory(c *gin.Context) {
	var category Models.Category
	id := c.Params.ByName("id")
	err := Models.ShowCategory(&category, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, category)
	}
	c.BindJSON(&category)
	err = Models.UpdateCategory(&category, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, category)
	} else {
		ApiHelpers.RespondJSON(c, 200, category)
	}
}

func DeleteCategory(c *gin.Context) {
	var category Models.Category
	id := c.Params.ByName("id")
	err := Models.DeleteCategory(&category, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, category)
	} else {
		ApiHelpers.RespondJSON(c, 200, category)
	}
}
