package main

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"

	"github.com/AMYMEME/board-golang/api"
	"github.com/AMYMEME/board-golang/config"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

var logger = config.GetLogger()

func main() {
	os.Setenv("PORT", "8090")
	r := setupRouter()
	if err := r.Run(); err != nil {
		logger.Errorf(errors.Wrap(err, "Fail gin engine start").Error())
	} //localhost:8080

	defer api.DB.MyDB.Close()
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           1 * time.Hour,
	}))

	r.GET("/members", api.GetAllMembers)
	r.GET("/member/:id", api.GetMember)
	r.DELETE("/member/:id", api.DeleteMember)
	r.PUT("/member/:id", api.UpdateMember)

	r.GET("/posts", api.GetAllPosts)
	r.POST("/post", api.AddPost)
	r.GET("/post/:id", api.GetPost)
	r.GET("/post/:id/comments", api.GetAllCommentsByPostId)
	r.POST("/post/:id/comment", api.AddComment)
	r.DELETE("/post/:id", api.DeletePost)
	r.PUT("/post/:id", api.UpdatePost)

	r.DELETE("/comment/:id", api.DeleteComment)
	r.PATCH("/comment/:id", api.UpdateComment)
	//댓글은 내용밖에 수정할 것이 없음

	//for auth
	r.POST("/auth/google", api.AuthGoogle)
	r.POST("/auth/naver", api.AuthNaver)
	r.POST("/auth/kakao", api.AuthKakao)

	return r
}
