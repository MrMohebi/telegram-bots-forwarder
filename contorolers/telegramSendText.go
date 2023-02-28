package contorolers

import (
	"github.com/gin-gonic/gin"
	tele "gopkg.in/telebot.v3"
	"net/http"
	"telegram-bots-forwarder/common"
	"time"
)

type data struct {
	Token  string `json:"token"`
	Text   string `json:"text"`
	ChatId int64  `json:"chatId,string"`
}

func TelegramSendText() gin.HandlerFunc {
	return func(context *gin.Context) {
		var requestBody data
		err := context.BindJSON(&requestBody)

		pref := tele.Settings{
			Token:  requestBody.Token,
			Poller: &tele.LongPoller{Timeout: 10 * time.Second},
		}

		telegram, err := tele.NewBot(pref)
		common.IsErr(err, "Could not create telegram bot")

		chat, err := telegram.ChatByID(requestBody.ChatId)
		common.IsErr(err)

		_, err = telegram.Send(chat, requestBody.Text)
		common.IsErr(err)

		context.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	}
}
