package telegram

import (
	"bytes"
	"encoding/json"
	tele "gopkg.in/telebot.v3"
	"log"
	"net/http"
	"telegram-bots-forwarder/common"
	"time"
)

type SendTelegramRequest struct {
	Message *tele.Message `json:"message"`
	User    *tele.User    `json:"user"`
}

func ListenOnTelegramBot(botName string, token string, callUrl string) {
	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle(tele.OnText, func(c tele.Context) error {
		data := SendTelegramRequest{
			Message: c.Message(),
			User:    c.Sender(),
		}

		dataJSON, err := json.Marshal(data)
		_, err = http.Post(callUrl, "application/json", bytes.NewBuffer(dataJSON))
		common.IsErr(err, "Could not make POST request to httpbin")

		return nil
	})

	println("bot " + botName + " is up and running")
	b.Start()

}
