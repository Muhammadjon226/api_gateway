package models

// Post ...
type Post struct {
	ID     int64  `json:"id"`
	UserID int64  `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	CreatedAt string `json:"created_at"`
	UpdatedAT string `json:"updated_at"`
}

 // PostRequest ...
type PostRequest struct {
	UserID int64  `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
// SuccessfullResponse ...
type SuccessfullResponse struct {
	Message string `json:"message"`
}

// ListPosts ...
type ListPosts struct {
	Posts []*Post `json:"post"`
	Count int64
}
