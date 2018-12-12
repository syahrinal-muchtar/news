package Controllers

import (
	"News/ApiHelpers"
	"News/Models"

	"github.com/gin-gonic/gin"
)

func AddPostTag(c *gin.Context) {
	var postTag Models.Post_Tag
	c.BindJSON(&postTag)

	err := Models.AddPostTag(&postTag, postTag.TagID, postTag.PostID)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, err)
	} else {
		ApiHelpers.RespondJSON(c, 200, postTag)
	}
}

//Get Post By Tag
func GetPost(c *gin.Context) {
	id := c.Params.ByName("id")
	var postTag []Models.PostbyTag
	err := Models.GetPost(&postTag, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, err)
	} else {
		ApiHelpers.RespondJSON(c, 200, postTag)
	}
}
