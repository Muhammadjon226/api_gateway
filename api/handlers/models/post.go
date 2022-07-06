package models

// Post ...
type Post struct {
	Id     int64  `json:"id"`
	UserId int64  `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// SuccessfullResponse ...
type SuccessfullResponse struct {
	Message string `json:"message"`
}

type UpdatePost struct {
	UserId int64  `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// ListPosts ...
type ListPosts struct {
	Posts []*Post `json:"post"`
	Count int64
}
