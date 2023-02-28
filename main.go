package main

import (
	"github.com/gin-gonic/gin"
	"telegram-bots-forwarder/common"
	"telegram-bots-forwarder/configs"
	"telegram-bots-forwarder/router"
	"telegram-bots-forwarder/telegram"
)

// nodemon --exec go run main.go --signal SIGTERM

func main() {
	configs.Setup()

	server := gin.Default()

	router.Routs(server)

	secs := configs.IniData.SectionStrings()
	for _, v := range secs {
		if v != "DEFAULT" {
			go telegram.ListenOnTelegramBot(v, configs.IniGet(v, "TOKEN"), configs.IniGet(v, "URL"))
		}
	}

	err := server.Run("localhost:8005")
	common.IsErr(err, "Err in starting server")

}
