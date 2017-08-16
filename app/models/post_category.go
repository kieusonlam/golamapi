package models

type (
	// PostCategory hold post category details
	PostCategory struct {
		ID          int
		Name        string
		Description string
		Posts       []Post `pg:"many_to_many:posts_categories"`
	}
)
