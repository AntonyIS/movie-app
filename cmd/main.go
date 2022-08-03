package main

import (
	"log"
	"net/http"
	"os"

	commentHandler "github.com/AntonyIS/movie-app/api"
	"github.com/AntonyIS/movie-app/app"
	redisCache "github.com/AntonyIS/movie-app/repository/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	commentRepo := commentRepo()
	commentService := app.NewcommentService(commentRepo)
	commentHandler := commentHandler.NewCommentHandler(commentService)

	router := gin.Default()

	router.POST("/comments/", commentHandler.CreateComment)
	router.GET("/comments/", commentHandler.GetComments)
	router.GET("/comments/:id", commentHandler.GetComment)
	router.PUT("/comments/:id", commentHandler.UpdateComment)
	router.DELETE("/comments/:id", commentHandler.DeleteComment)
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"message": "Welcome to star wars API",
		})
	})

	router.Run(port())
}

func port() string {
	port := ":8000"

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return port
}

func commentRepo() app.CommentRepository {
	repo, err := redisCache.NewRedisRepository("redis://localhost:6379")

	if err != nil {
		log.Fatal("redis server not connected: ", err)
		return nil
	}
	return repo
}
