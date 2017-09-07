package models

import "aahframework.org/log.v0"

// PostCategory Post Category relation
type PostCategory struct {
	ID         int `json:"id"`
	PostID     int `json:"post_id"`
	CategoryID int `json:"category_id"`
}

// CreatePostCatRelation add relation between post and category
func CreatePostCatRelation(postcat *PostCategory) (*PostCategory, error) {
	// err := db.Insert(postcat)
	create, err := db.Model(postcat).
		Column("id").
		Where("post_id = ?post_id").
		Where("category_id = ?category_id").
		OnConflict("DO NOTHING").
		Returning("id").
		SelectOrInsert()
	if err != nil {
		return nil, err
	}
	return postcat, nil
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
func DelPostCatRel(id int) (int, error) {
	err := db.Delete(&PostCategory{
		ID: id,
	})
	if err != nil {
		return 0, err
	}
	return id, nil
}
