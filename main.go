package main

import (
	"github.com/gin-gonic/gin"
	"telegram-bots-forwarder/common"
	"telegram-bots-forwarder/configs"
	"telegram-bots-forwarder/router"
)

// nodemon --exec go run main.go --signal SIGTERM

func main() {
	configs.Setup()

	server := gin.Default()

	router.Routs(server)

	err := server.Run("localhost:8005")
	common.IsErr(err, "Err in starting server")
}
