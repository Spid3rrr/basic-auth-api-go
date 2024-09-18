package entities

import (
	notes "github.com/Spid3rrr/basic-auth-api/entities/notes"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	notes.SetupRoutes(router)
}
