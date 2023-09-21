package handler

import (
	"codehub.devcloud.cn-north-4.huaweicloud.com/taqse2020_mihoyo00001/blog_backend/model"
	"codehub.devcloud.cn-north-4.huaweicloud.com/taqse2020_mihoyo00001/blog_backend/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func LoginHandler(c *gin.Context) {
	// 绑定客户端提交的请求参数
	var req model.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errcode": 400, "errmsg": "bad request"})
		return
	}

	// 调用 service 提供的 Login 函数进行用户登录
	log.Info("Login Req : %v", req)
	token, err := service.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errcode": 500, "errmsg": err.Error()})
		return
	}

	// 登录成功，返回 token
	c.JSON(http.StatusOK, gin.H{"errcode": 0, "errmsg": "success", "data": model.LoginResp{Token: token}})
}
