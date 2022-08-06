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

	movieRepo := movieRepo()
	MovieService := app.NewMovieService(movieRepo)
	movieHandler := handler.NewMovieHandler(MovieService)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"message": "Welcome",
		})
	})
	router.GET("/characters/seed", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"message": seed.SeedCharacters(),
		})
	})

	router.GET("/movies/seed", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"message": seed.SeedMovies(),
		})
	})

	router.POST("/comments/:movie_id/", commentHandler.CreateComment)
	router.GET("/comments/:movie_id", commentHandler.GetComments)
	router.GET("/comments/:movie_id/:comment_id", commentHandler.GetComment)
	router.PUT("/comments/:movie_id/:comment_id", commentHandler.UpdateComment)
	router.DELETE("/comments/:movie_id/:comment_id", commentHandler.DeleteComment)
	router.POST("/characters/", characterHandler.CreateCharacter)
	router.GET("/characters/", characterHandler.GetCharacters)
	router.GET("/characters/:id", characterHandler.GetCharacter)
	router.PUT("/characters/:id", characterHandler.UpdateCharacter)
	router.DELETE("/characters/:id", characterHandler.DeleteCharacter)
	router.POST("/movies/", movieHandler.CreateMovie)
	router.GET("/movies/", movieHandler.GetMovies)
	router.GET("/movies/:id", movieHandler.GetMovie)
	router.GET("/movies/comments/:id", movieHandler.GetMovieComments)
	router.GET("/movies/characters/:id", movieHandler.GetMovieCharacters)
	router.PUT("/movies/:id", movieHandler.UpdateMovie)
	router.DELETE("/movies/:id", movieHandler.DeleteMovie)
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
	repo, err := redisCache.NewCharacterRedisRepository("redis://localhost:6379")

	if err != nil {
		log.Fatal("redis server not connected: ", err)
		return nil
	}
	return repo
}

func movieRepo() app.MovieRepository {
	repo, err := redisCache.NewMovieRedisRepository("redis://localhost:6379")

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
