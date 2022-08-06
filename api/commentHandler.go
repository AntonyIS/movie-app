package api

import (
	"net/http"

	"github.com/AntonyIS/movie-app/app"
	"github.com/gin-gonic/gin"
)

type CommentHandler interface {
	CreateComment(*gin.Context)
	GetComment(*gin.Context)
	GetComments(*gin.Context)
	UpdateComment(*gin.Context)
	DeleteComment(*gin.Context)
}

type commentHandler struct {
	commentService app.CommentService
}

func NewCommentHandler(commentService app.CommentService) CommentHandler {
	return &commentHandler{
		commentService,
	}
}

func (h *commentHandler) CreateComment(c *gin.Context) {
	movie_id := c.Param("movie_id")
	var comment app.Comment

	comment.MovieID = movie_id
	comment.URL = c.Request.Host + c.Request.URL.Path

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	_, err := h.commentService.CreateComment(&comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"comemnts": comment,
	})
}

func (h *commentHandler) GetComments(c *gin.Context) {
	comments, err := h.commentService.GetComments()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"comments": comments,
	})
}

func (h *commentHandler) GetComment(c *gin.Context) {
	id := c.Param("id")
	comment, err := h.commentService.GetComment(id)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"comment": err,
		})
		return
	}
	c.JSON(http.StatusCreated, comment)
}

func (h *commentHandler) UpdateComment(c *gin.Context) {
	var comment app.Comment

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": app.ErrorInvalidItem,
		})
		return
	}

	data, err := h.commentService.UpdateComment(&comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"comments": data,
	})
}

func (h *commentHandler) DeleteComment(c *gin.Context) {
	commentID := c.Param("id")
	if err := h.commentService.DeleteComment(commentID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": app.ErrorInvalidItem,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "comment deleted successfuly",
	})

}
