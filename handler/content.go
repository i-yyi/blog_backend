package handler

import (
	"codehub.devcloud.cn-north-4.huaweicloud.com/taqse2020_mihoyo00001/blog_backend/model"
	"codehub.devcloud.cn-north-4.huaweicloud.com/taqse2020_mihoyo00001/blog_backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateCommentHandler(c *gin.Context) {
	// 绑定客户端提交的请求参数
	var req model.CreateCommentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errcode": 400, "errmsg": "bad request"})
		return
	}

	// 调用 service 提供的 CreateComment 函数进行添加评论
	if err := service.CreateComment(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errcode": 500, "errmsg": err.Error()})
		return
	}

	// 添加评论成功，返回结果
	c.JSON(http.StatusOK, gin.H{"errcode": 0, "errmsg": "success", "data": model.CreateCommentResp{}})
}
func CreatePostHandler(c *gin.Context) {
	// 绑定客户端提交的请求参数
	var req model.CreatePostReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errcode": 400, "errmsg": "bad request"})
		return
	}

	// 调用 service 提供的 CreatePost 函数进行文章发表
	if err := service.CreatePost(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errcode": 500, "errmsg": err.Error()})
		return
	}

	// 文章发表成功，返回结果
	c.JSON(http.StatusOK, gin.H{"errcode": 0, "errmsg": "success", "data": model.CreatePostResp{}})
}

func ListPostHandler(c *gin.Context) {
	// 调用 service 提供的 GetPosts 函数获取文章列表
	posts, err := service.ListPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errcode": 500, "errmsg": err.Error()})
		return
	}

	// 获取文章列表成功，返回结果
	c.JSON(http.StatusOK, gin.H{"errcode": 0, "errmsg": "success", "data": model.ListPostResp{Posts: posts}})
}

func PostDetailHandler(c *gin.Context) {
	// 从 request 中获取 postID
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errcode": 400, "errmsg": "bad request"})
		return
	}

	// 调用 service 提供的 GetPost 函数获取指定文章的详细信息
	post, err := service.ViewPostDetail(postID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"errcode": 404, "errmsg": "post not found"})
		return
	}

	// 获取指定文章详细信息成功，返回结果
	c.JSON(http.StatusOK, gin.H{"errcode": 0, "errmsg": "success", "data": post})
}

func DeletePostHandler(c *gin.Context) {
	// 从 request 中获取 postID
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errcode": 400, "errmsg": "bad request"})
		return
	}

	// 调用 service 提供的 DeletePost 函数删除指定文章
	if err := service.DeletePost(postID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errcode": 500, "errmsg": err.Error()})
		return
	}

	// 删除文章成功，返回结果
	c.JSON(http.StatusOK, gin.H{"errcode": 0, "errmsg": "success", "data": model.DeletePostResp{}})
}
