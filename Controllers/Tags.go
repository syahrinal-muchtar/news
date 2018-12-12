package Controllers

import (
	"News/ApiHelpers"
	"News/Models"

	"github.com/gin-gonic/gin"
)

func ShowTags(c *gin.Context) {
	var tag []Models.Tag
	err := Models.ShowTags(&tag)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, tag)
	} else {
		ApiHelpers.RespondJSON(c, 200, tag)
	}
}

func ShowTag(c *gin.Context) {
	id := c.Params.ByName("id")
	var tag Models.Tag
	err := Models.ShowTag(&tag, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, tag)
	} else {
		ApiHelpers.RespondJSON(c, 200, tag)
	}
}

func AddTag(c *gin.Context) {
	var tag Models.Tag
	c.BindJSON(&tag)
	err := Models.AddTag(&tag)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, tag)
	} else {
		ApiHelpers.RespondJSON(c, 200, tag)
	}
}

func UpdateTag(c *gin.Context) {
	var tag Models.Tag
	id := c.Params.ByName("id")
	err := Models.ShowTag(&tag, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, tag)
	}
	c.BindJSON(&tag)
	err = Models.UpdateTag(&tag, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, tag)
	} else {
		ApiHelpers.RespondJSON(c, 200, tag)
	}
}

func DeleteTag(c *gin.Context) {
	var tag Models.Tag
	id := c.Params.ByName("id")
	err := Models.DeleteTag(&tag, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, tag)
	} else {
		ApiHelpers.RespondJSON(c, 200, tag)
	}
}
