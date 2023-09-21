package model

import "time"

type BaseResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}
type PostAbstract struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Abstract   string    `json:"abstract"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
type LoginResp struct {
	BaseResp
	Token string `json:"token"`
}

type CreatePostResp struct {
	BaseResp
}
type ListPostResp struct {
	BaseResp
	Posts []PostAbstract `json:"posts"`
}
type PostDetailResp struct {
	BaseResp
	Post
	Comments []Comment `json:"comments"`
}
type DeletePostResp struct {
	BaseResp
}
type CreateCommentResp struct {
	BaseResp
}
type ListCommentResp struct {
	BaseResp
	Comments []Comment `json:"comments"`
}
