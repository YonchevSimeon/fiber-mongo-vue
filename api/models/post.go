package models

import "github.com/Kamva/mgm/v2"

type Post struct {
	mgm.DefaultModel `bson:",inline"`
	UserID           string `json:"userId" bson:"userId,omitempty"`
	Title            string `json:"title" bson:"title"`
	Content          string `json:"content" bson:"content"`
}

func CreatePost(userId, title, content string) *Post {
	return &Post{
		UserID:  userId,
		Title:   title,
		Content: content,
	}
}
