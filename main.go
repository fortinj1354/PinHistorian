package main

import (
	"github.com/fortinj1354/Pin-Historian/app"
	"github.com/fortinj1354/Pin-Historian/models"
	"github.com/fortinj1354/Pin-Historian/settings"
	"github.com/getsentry/raven-go"
	"github.com/gin-contrib/sentry"
	"github.com/gin-gonic/gin"
)

func main() {
	settings.LoadSettings()
	router := gin.Default()
	models.MakeDB(settings.GetDatabaseName())

	router.Use(sentry.Recovery(raven.DefaultClient, false))

	router.POST("/", app.HandlePost)
	router.GET("/health", app.HandleHealth)
	router.GET("teams/:teamId/pins", app.HandleGetPins)
	router.GET("/teams/:teamId/channels", app.HandleGetChannels)
	router.GET("/teams/:teamId/channels/:channelId/pins", app.HandleGetPins)

	router.Run()
}
