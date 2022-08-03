package api

import (
	"net/http"

	"github.com/AntonyIS/movie-app/app"
	"github.com/gin-gonic/gin"
)

type CharacterHandler interface {
	CreateCharacter(*gin.Context)
	GetCharacter(*gin.Context)
	GetCharacters(*gin.Context)
	UpdateCharacter(*gin.Context)
	DeleteCharacter(*gin.Context)
}

type characterHandler struct {
	characterService app.CharacterService
}

func NewCharacterHandler(characterService app.CharacterService) CharacterHandler {
	return &characterHandler{
		characterService,
	}
}

func (h *characterHandler) CreateCharacter(c *gin.Context) {
	var Character app.Character

	if err := c.ShouldBindJSON(&Character); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	_, err := h.characterService.CreateCharacter(&Character)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"Character": Character,
	})
}

func (h *characterHandler) GetCharacters(c *gin.Context) {
	Characters, err := h.characterService.GetCharacters()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"Characters": Characters,
	})
}

func (h *characterHandler) GetCharacter(c *gin.Context) {
	id := c.Param("id")
	Character, err := h.characterService.GetCharacter(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"Characters": Character,
	})
}

func (h *characterHandler) UpdateCharacter(c *gin.Context) {
	var character app.Character

	if err := c.ShouldBindJSON(&character); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": app.ErrorInvalidItem,
		})
		return
	}

	data, err := h.characterService.UpdateCharacter(&character)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"Characters": data,
	})
}

func (h *characterHandler) DeleteCharacter(c *gin.Context) {
	CharacterID := c.Param("id")
	if err := h.characterService.DeleteCharacter(CharacterID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": app.ErrorInvalidItem,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Character deleted successfuly",
	})

}
