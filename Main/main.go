package main

import (
	injector "github.com/dilgeto/borrador-optica-backend/Main/Injector"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default
	injector.InjectDependencies(router)
	router.Run("localhost:8080")
}
