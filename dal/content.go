package dal

import "codehub.devcloud.cn-north-4.huaweicloud.com/taqse2020_mihoyo00001/blog_backend/model"

// GetPostAbstract 获取文章摘要
func GetPostAbstract(content string, length int) string {
	if length >= len(content) {
		return content
	}
	runes := []rune(content)
	return string(runes[:length])
}
func GetPostById(id int) (*model.Post, error) {
	post := &model.Post{}
	result := db.First(post, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return post, nil
}
func DeletePost(post *model.Post) error {
	result := db.Delete(post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func GetCommentById(id int) (*model.Comment, error) {
	comment := &model.Comment{}
	result := db.First(comment, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return comment, nil
}
func ListCommentsByPostId(postId int) ([]model.Comment, error) {
	var comments []model.Comment
	result := db.Where("post_id = ?", postId).Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}
	return comments, nil
}
func CreateComment(comment *model.Comment) error {
	result := db.Create(comment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func CreatePost(post *model.Post) error {
	result := db.Create(post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func ListPosts() ([]model.PostAbstract, error) {
	posts := []model.Post{}
	result := db.Select("id, title, SUBSTR(content, 1, 100) as abstract, create_time, update_time").Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}

	// 封装响应参数
	postAbstracts := []model.PostAbstract{}
	for _, post := range posts {
		postAbstract := model.PostAbstract{
			Id:         int(post.ID),
			Title:      post.Title,
			Abstract:   GetPostAbstract(post.Content, 100),
			CreateTime: post.CreateTime,
			UpdateTime: post.UpdateTime,
		}
		postAbstracts = append(postAbstracts, postAbstract)
	}
	return postAbstracts, nil
}
