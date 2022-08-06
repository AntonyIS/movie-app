package api

import (
	"net/http"

	"github.com/AntonyIS/movie-app/app"
	"github.com/gin-gonic/gin"
)

type MovieHandler interface {
	CreateMovie(*gin.Context)
	GetMovie(*gin.Context)
	GetMovies(*gin.Context)
	UpdateMovie(*gin.Context)
	DeleteMovie(*gin.Context)
	GetMovieComments(*gin.Context)
	GetMovieCharacters(*gin.Context)
}

type movieHandler struct {
	movieService app.MovieService
}

func NewMovieHandler(movieService app.MovieService) MovieHandler {
	return &movieHandler{
		movieService,
	}
}

func (h *movieHandler) CreateMovie(c *gin.Context) {
	var movie app.Movie

	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	_, err := h.movieService.CreateMovie(&movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"movie": movie,
	})
}

func (h *movieHandler) GetMovies(c *gin.Context) {
	movies, err := h.movieService.GetMovies()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"movies": movies,
	})
}

func (h *movieHandler) GetMovie(c *gin.Context) {
	id := c.Param("id")
	movie, err := h.movieService.GetMovie(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"movies": movie,
	})
}

func (h *movieHandler) UpdateMovie(c *gin.Context) {
	var movie app.Movie

	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": app.ErrorInvalidItem,
		})
		return
	}

	data, err := h.movieService.UpdateMovie(&movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"movies": data,
	})
}

func (h *movieHandler) DeleteMovie(c *gin.Context) {
	movieID := c.Param("id")
	if err := h.movieService.DeleteMovie(movieID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": app.ErrorInvalidItem,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Movie deleted successfuly",
	})

}

func (h *movieHandler) GetMovieComments(c *gin.Context) {
	id := c.Param("id")
	comments, err := h.movieService.GetMovieComments(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"comments": comments,
	})
}
func (h *movieHandler) GetMovieCharacters(c *gin.Context) {
	id := c.Param("id")
	characters, err := h.movieService.GetMovieCharacters(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, characters)
}
