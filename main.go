package main

import (
	"os"

	"github.com/AMYMEME/board-golang/config"
	"github.com/AMYMEME/board-golang/controller"
	"github.com/gin-contrib/cors"
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

	defer controller.DB.MyDB.Close()
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default()) //All CORS permit

	r.GET("/members", controller.GetAllMembers)
	r.POST("/member", controller.AddMember)
	r.GET("/member/:id", controller.GetMember)
	r.DELETE("/member/:id", controller.DeleteMember)
	r.PUT("/member/:id", controller.UpdateMember)

	r.GET("/posts", controller.GetAllPosts)
	r.POST("/post", controller.AddPost)
	r.GET("/post/:id", controller.GetPost)
	r.GET("/post/:id/comments", controller.GetAllCommentsByPostId)
	r.POST("/post/:id/comment", controller.AddComment)
	r.DELETE("/post/:id", controller.DeletePost)
	r.PUT("/post/:id", controller.UpdatePost)

	r.DELETE("/comment/:id", controller.DeleteComment)
	r.PATCH("/comment/:id", controller.UpdateComment)
	//댓글은 내용밖에 수정할 것이 없음

	return r
}
