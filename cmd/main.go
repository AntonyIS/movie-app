package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	handler "github.com/AntonyIS/movie-app/api"
	"github.com/AntonyIS/movie-app/app"
	postgres "github.com/AntonyIS/movie-app/repository/postgresdb"
	redisCache "github.com/AntonyIS/movie-app/repository/redis"
	seed "github.com/AntonyIS/movie-app/seed"
	"github.com/gin-gonic/gin"
)

func main() {
	commentRepo := commentRepo()
	commentService := app.NewcommentService(commentRepo)
	commentHandler := handler.NewCommentHandler(commentService)

	characterRepo := characterRepo()
	characterService := app.NewCharacterService(characterRepo)
	characterHandler := handler.NewCharacterHandler(characterService)

	seed.PostCharacters()
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"message": "Welcome to star wars API",
		})
	})
	router.POST("/comments/", commentHandler.CreateComment)
	router.GET("/comments/", commentHandler.GetComments)
	router.GET("/comments/:id", commentHandler.GetComment)
	router.PUT("/comments/:id", commentHandler.UpdateComment)
	router.DELETE("/comments/:id", commentHandler.DeleteComment)
	router.POST("/characters/", characterHandler.CreateCharacter)
	router.GET("/characters/", characterHandler.GetCharacters)
	router.GET("/characters/:id", characterHandler.GetCharacter)
	router.PUT("/characters/:id", characterHandler.UpdateCharacter)
	router.DELETE("/characters/:id", characterHandler.DeleteCharacter)
	router.Run(port())
}

func port() string {
	port := ":8000"

	if os.Getenv("SERVERPORT") != "" {
		port = fmt.Sprintf(":%s", os.Getenv("SERVERPORT"))
	}
	return port
}

func characterRepo() app.CharacterRepository {
	repo, err := redisCache.NewRedisRepository("redis://localhost:6379")

	if err != nil {
		log.Fatal("redis server not connected: ", err)
		return nil
	}
	return repo
}

func commentRepo() app.CommentRepository {
	repo := postgres.NewRepostory()
	return repo
}
