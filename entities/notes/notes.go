package notes

import (
	"net/http"
	"strconv"

	auth "github.com/Spid3rrr/basic-auth-api/auth"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Note struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Body   string `json:"body" validate:"required,min=6,max=256"`
}

var Notes = []Note{}

func getNoteByID(c *gin.Context) {
	id := c.Param("id")
	noteID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}
	for _, note := range Notes {
		if note.ID == noteID {
			c.JSON(http.StatusOK, note)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Note not found"})
}

func addNote(c *gin.Context) {
	var note Note
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note.ID = len(Notes) + 1
	note.Author = c.MustGet("username").(string)

	Notes = append(Notes, note)
	c.JSON(http.StatusCreated, note)
}

func deleteNoteByID(c *gin.Context) {
	id := c.Param("id")
	noteID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}
	for i, note := range Notes {
		if note.ID == noteID {
			Notes = append(Notes[:i], Notes[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Note deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Note not found"})
}

func SetupRoutes(router *gin.Engine) {
	notes := router.Group("/notes")
	{
		notes.GET("/:id", auth.CheckAuth, checkAuthor, getNoteByID)
		notes.POST("", auth.CheckAuth, addNote)
		notes.DELETE("/:id", deleteNoteByID)
	}
}
