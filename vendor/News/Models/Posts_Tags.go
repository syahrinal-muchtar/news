package Models

import (
	"News/Config"

	_ "github.com/go-sql-driver/mysql"
)

func AddPostTag(b *Post_Tag, tagId uint, postId uint) (err error) {
	Config.DB.Where(&Post_Tag{PostID: postId, TagID: tagId}).First(b)
	if err = Config.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

//Get Post By Tag
func GetPost(b *[]PostbyTag, id string) (err error) {
	if err := Config.DB.Raw("SELECT tags.`name` AS tagname, posts.title, posts.content, users.`name` AS author FROM post_tag INNER JOIN posts ON post_tag.post_id = posts.id INNER JOIN tags ON post_tag.tag_id = tags.id INNER JOIN users ON posts.user_id = users.id WHERE post_tag.deleted_at IS NULL AND tags.id = ?", id).Scan(b).Error; err != nil {
		return err
	}
	return nil
}
