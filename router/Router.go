package router

import (
	"github.com/gin-gonic/gin"
	"telegram-bots-forwarder/contorolers"
)

func Routs(r *gin.Engine) {
	AssetsRoute(r)
	r.LoadHTMLGlob("templates/**/*.html")

	r.GET("/", contorolers.Index())
	r.GET("/docs", contorolers.Docs())

	telegram := r.Group("/telegram")
	{
		telegram.POST("/sendText", contorolers.TelegramSendText())
	}

}
