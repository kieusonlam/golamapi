package models

import "aahframework.org/log.v0"

type (
	// PostsCategories Post Category relation
	PostsCategories struct {
		ID         int `json:"id"`
		PostID     int `json:"post_id"`
		CategoryID int `json:"category_id"`
	}
)

// PostCatRelation add relation between post and category
func PostCatRelation(postid int, catid int) *PostsCategories {
	postcat := &PostsCategories{
		PostID:     postid,
		CategoryID: catid,
	}
	err := db.Insert(postcat)
	if err != nil {
		log.Error(err)
	}
	return postcat
}

// GetPostCatRelations use to get all posts.
func GetPostCatRelations() []PostsCategories {
	var postcatrels []PostsCategories
	err := db.Model(&postcatrels).Select()
	if err != nil {
		log.Error(err)
	}
	return postcatrels
}
