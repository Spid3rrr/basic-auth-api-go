package main

import (
	"os"

	auth "github.com/Spid3rrr/basic-auth-api/auth"
	entities "github.com/Spid3rrr/basic-auth-api/entities"
	"github.com/gin-gonic/gin"
)

func main() {

	os.Setenv("DEFAULT_USERNAME", "admin")
	os.Setenv("DEFAULT_PASSWORD", "admin")
	os.Setenv("JWT_SECRET", "secret")

	router := gin.Default()
	entities.SetupRoutes(router)
	auth.SetupRoutes(router)

	router.Run(":8080")
}
