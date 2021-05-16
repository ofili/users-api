package app

import (
	"github.com/gin-gonic/gin"
	"github.com/ofili/bookstore_users-api/logger"

)

var (
	router = gin.Default()
)

func StartAppication() {
	mapURLs()

	logger.Info("about to start application...")
	router.Run(":5000")
}
