package main

import (
	"thinkdecideact/src/config"
	"thinkdecideact/src/routers"
	"thinkdecideact/src/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config := config.GetConfig()
	DB := utils.GetDB(config.DB)

	apiGroup := r.Group("/api")
	{
		routers.ApiRouter(apiGroup, DB)
	}

	// By default it serves on :8080 unless a PORT environment variable was defined.
	// router.Run(":3000") for a hard coded port
	r.Run(":8080")
}
