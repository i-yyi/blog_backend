package model

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreatePostReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ListPostReq struct {
}

type PostDetailReq struct {
	Id int `json:"id"`
}
type DeletePostReq struct {
	Id int `json:"id"`
}
type CreateCommentReq struct {
	PostId  int    `json:"post_id"`
	Content string `json:"content"`
}
type ListCommentReq struct {
	Id int `json:"id"`
}
