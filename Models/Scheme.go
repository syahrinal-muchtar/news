package Models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Age      string `json:"age"`
	Address  string `json:"address"`
	Nation   string `json:"nation"`
	Province string `json:"province"`
	City     string `json:"city"`
	Postcode string `json:"postcode"`
	Role     string `json:"role"`
}

type Comment struct {
	gorm.Model
	Content string `json:"content"`
	UserID  uint   `gorm:"size:10"`
	User    User
}

type Post struct {
	gorm.Model
	Title      string `json:"title"`
	Content    string `json:"content"`
	UserID     uint   `gorm:"size:10" json:"user_id"`
	User       User
	CategoryID uint `gorm:"size:10" json:"category_id"`
	Category   Category
	Picture    string
}

type News struct {
	ID      int
	Title   string
	Date    time.Time
	Author  string
	Image   string
	Content string
}

type HotNews struct {
	ID      int
	Title   string
	Date    string
	Author  string
	Image   string
	Content string
}

type Category struct {
	gorm.Model
	Name string `json:"name"`
}

type Tag struct {
	gorm.Model
	Name string `json:"name"`
}

type Post_Tag struct {
	gorm.Model
	PostID uint `gorm:"size:10" json:"post_id"`
	Post   Post
	TagID  uint `gorm:"size:10" json:"tag_id"`
	Tag    Tag
}

type PostbyTag struct {
	Title   string
	Author  string
	Content string
	Tagname string
}

func (b *User) TableName() string {
	return "users"
}

func (b *Post) TableName() string {
	return "posts"
}

func (b *Tag) TableName() string {
	return "tags"
}

func (b *Comment) TableName() string {
	return "comments"
}

func (b *Category) TableName() string {
	return "categories"
}

func (b *Post_Tag) TableName() string {
	return "post_tag"
}

func (b *News) TableName() string {
	return "news"
}
