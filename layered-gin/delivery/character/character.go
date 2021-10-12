// layered-gin/delivery/character/character.go

package character

import (
	"encoding/json"
	"io/ioutil"
	"layered-gin/entity"
	"layered-gin/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CharacterHandler struct {
	usecase usecase.Character
}

func New(character usecase.Character) CharacterHandler {
	return CharacterHandler{usecase: character}
}

// GetAll() makes a GET request and returns all characters
func (c CharacterHandler) GetAll(g *gin.Context) {
	var characters []entity.Character

	chars, err := c.usecase.GetAll(characters)
	if err != nil {
		g.JSON(500, gin.H{
			"status": http.StatusInternalServerError,
			"error": err,
			"msg": "unable to get all characters",
		})
		g.Abort()
		return
	}

		g.JSON(200, gin.H{
			"status": http.StatusOK,
			"characters": chars,
		})
}

// GetByID() makes a GET request and returns a specific character
func (c CharacterHandler) GetByID(g *gin.Context) {
	params := g.Param("id")
	id, err := strconv.ParseUint(params, 10, 32)
	if err != nil {
		g.JSON(400, gin.H{
			"status": http.StatusBadRequest,
			"error": err,
			"msg": "unable to parse parameter",
		})
		g.Abort()
		return
	}

	char, err := c.usecase.GetByID(uint(id))
	if err != nil {
		g.JSON(400, gin.H{
			"status": http.StatusBadRequest,
			"error": err,
			"msg": "unable to get character",
		})
		g.Abort()
		return
	}

	g.JSON(200, gin.H{
		"status": http.StatusOK,
		"character": char,
	})
}

// Create() makes a POST request and returns a newly created character
func (c CharacterHandler) Create(g *gin.Context) {
	var character entity.Character

	body, err := ioutil.ReadAll(g.Request.Body)
	if err != nil {
		g.JSON(500, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error": err,
			"msg": "unable to process request body",
		})
		g.Abort()
		return
	}

	char, err := c.usecase.Create(character)
	if err != nil {
		g.JSON(400, gin.H{
			"status": http.StatusBadRequest,
			"error": err,
			"msg": "unable to create character",
		})
		g.Abort()
		return
	}

	err = json.Unmarshal(body, &char)
	if err != nil {
		g.JSON(500, gin.H{
			"status": http.StatusUnprocessableEntity, 
			"error": err,
			"msg": "unable to unmarshal character",
		})
		g.Abort()
		return
	}

	entity.Prepare(character)
	err = entity.Validate("create", character)
	if err != nil {
		g.JSON(500, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error": err,
			"msg": "unable to validate character",
		})
		g.Abort()
		return
	}

	char, err = c.usecase.Create(character)
	if err != nil {
		g.JSON(500, gin.H{
			"status": http.StatusInternalServerError,
			"error": err,
			"msg": "unable to create character",
		})
		g.Abort()
		return
	}
	
	g.JSON(200, gin.H{
		"status": http.StatusOK,
		"character": char,
	})
}