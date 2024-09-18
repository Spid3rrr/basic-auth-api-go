package notes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func checkAuthor(c *gin.Context) {
	username := c.MustGet("username").(string)

	id := c.Param("id")

	noteID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	for _, note := range Notes {
		if note.ID == noteID {
			if note.Author != username {
				c.AbortWithStatus(401)
				return
			}
			break
		}
	}

	c.Next()
}
