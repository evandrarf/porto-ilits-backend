package domain

var (
	GET_POSTS_SUCCESS = "Posts retrieved successfully"
	GET_POSTS_FAILED  = "Failed to retrieve posts"

	CREATE_POST_SUCCESS = "Post created successfully"
	CREATE_POST_FAILED  = "Failed to create post"
)

type (
	PostCreateRequest struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	Post struct {
		ID      uint   `json:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
		CreatedAt int64  `json:"created_at"`
	}

	PostResponse Post
)