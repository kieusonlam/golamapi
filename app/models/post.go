package models

type (
	// Post struct hold the post details
	Post struct {
		ID         int
		Title      string
		Content    string
		Caterories string
	}
)

// CreateNewPost use to create new post.
func CreateNewPost(title string, content string) *Post {
	post := &Post{
		Title:      title,
		Content:    content,
		Caterories: "1",
	}
	err := db.Insert(post)
	if err != nil {
		panic(err)
	}
	return post
}

// GetPostsData use to get all posts.
func GetPostsData() []Post {
	var posts []Post
	err := db.Model(&posts).Select()
	if err != nil {
		panic(err)
	}
	return posts
}
