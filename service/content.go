package service

import (
	"codehub.devcloud.cn-north-4.huaweicloud.com/taqse2020_mihoyo00001/blog_backend/dal"
	"codehub.devcloud.cn-north-4.huaweicloud.com/taqse2020_mihoyo00001/blog_backend/model"
	"errors"
)

func ListCommentsByPostId(postId int) ([]model.Comment, error) {
	comments, err := dal.ListCommentsByPostId(postId)
	if err != nil {
		// 处理获取评论列表失败的情况
		return nil, errors.New("get comments failed")
	}
	return comments, nil
}
func CreatePost(postReq *model.CreatePostReq) error {
	post := &model.Post{
		Title:   postReq.Title,
		Content: postReq.Content,
	}

	err := dal.CreatePost(post)
	if err != nil {
		// 处理创建文章失败的情况
		return errors.New("create post failed")
	}

	return nil
}
func CreateComment(commentReq *model.CreateCommentReq) error {
	comment := &model.Comment{
		PostID:  uint(commentReq.PostId),
		Content: commentReq.Content,
	}

	err := dal.CreateComment(comment)
	if err != nil {
		// 处理创建评论失败的情况
		return errors.New("create comment failed")
	}
	return nil
}
func ListPosts() ([]model.PostAbstract, error) {
	posts, err := dal.ListPosts()
	if err != nil {
		// 处理获取文章列表失败的情况
		return nil, errors.New("get posts failed")
	}
	return posts, nil
}
func ViewPostDetail(postId int) (*model.PostDetailResp, error) {
	// 获取文章信息
	post, err := dal.GetPostById(postId)
	if err != nil {
		// 处理获取文章失败的情况
		return nil, errors.New("get post failed")
	}

	// 获取评论列表
	comments, err := dal.ListCommentsByPostId(postId)
	if err != nil {
		// 处理获取评论列表失败的情况
		return nil, errors.New("get comments failed")
	}

	// 封装响应参数
	postDetail := &model.PostDetailResp{
		BaseResp: model.BaseResp{
			Errcode: 0,
			Errmsg:  "success",
		},
		Post: model.Post{
			ID:         post.ID,
			Title:      post.Title,
			Content:    post.Content,
			CreateTime: post.CreateTime,
			UpdateTime: post.UpdateTime,
		},
		Comments: comments,
	}
	return postDetail, nil
}

func DeletePost(id int) error {
	// 获取需要删除的文章
	post, err := dal.GetPostById(id)
	if err != nil {
		// 处理获取文章失败的情况
		return errors.New("get post failed")
	}

	// 删除文章
	err = dal.DeletePost(post)
	if err != nil {
		// 处理删除文章失败的情况
		return errors.New("delete post failed")
	}

	return nil
}
