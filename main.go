package main

import (
	"codehub.devcloud.cn-north-4.huaweicloud.com/taqse2020_mihoyo00001/blog_backend/conf"
	"codehub.devcloud.cn-north-4.huaweicloud.com/taqse2020_mihoyo00001/blog_backend/dal"
	"codehub.devcloud.cn-north-4.huaweicloud.com/taqse2020_mihoyo00001/blog_backend/handler"
	"github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
	log "github.com/sirupsen/logrus"
    "time"
)

func main() {
	config.InitConfig()
	dal.InitDB()
	r := gin.Default()
    
    r.Use(cors.New(cors.Config{
		AllowAllOrigins:     true,
		AllowMethods:     []string{"GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))
    r.POST("/posts", handler.CreatePostHandler)
	r.GET("/posts", handler.ListPostHandler)
    r.POST("/login", handler.LoginHandler)
	r.POST("/comments", handler.CreateCommentHandler)
    r.GET("/posts/:id", handler.PostDetailHandler)
	err := r.Run()
	if err != nil {
		log.Panic("[Gin] Error %v, Panic!", err)
	}
}
