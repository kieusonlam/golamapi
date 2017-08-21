package models

import "aahframework.org/log.v0"

// PostCategory Post Category relation
type PostCategory struct {
	ID         int `json:"id"`
	PostID     int `json:"post_id"`
	CategoryID int `json:"category_id"`
}

// PostCatRelation add relation between post and category
func PostCatRelation(postid int, catid int) *PostCategory {
	postcat := &PostCategory{
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
func GetPostCatRelations() []PostCategory {
	var postcatrels []PostCategory
	err := db.Model(&postcatrels).Select()
	if err != nil {
		log.Error(err)
	}
	return postcatrels
}

// DelPostCatRel remove relation record
func DelPostCatRel(id int) interface{} {
	err := db.Delete(&PostCategory{
		ID: id,
	})
	if err != nil {
		log.Error(err)
	}
	return id
}
