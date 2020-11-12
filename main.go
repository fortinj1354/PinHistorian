package main

import (
	"fmt"
	"github.com/fortinj1354/Pin-Historian/app"
	"github.com/fortinj1354/Pin-Historian/models"
	"github.com/fortinj1354/Pin-Historian/settings"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

func main() {
	settings.LoadSettings()
	router := gin.Default()
	models.MakeDB(settings.GetDatabaseName())

	if settings.GetSentryDSN() != "" && settings.GetSentryEnvironment() != "" {
		if err := sentry.Init(sentry.ClientOptions{
			Dsn:              settings.GetSentryDSN(),
			Environment:      settings.GetSentryEnvironment(),
			Debug:            true,
			AttachStacktrace: true,
		}); err != nil {
			fmt.Printf("Sentry initialization failed: %v\n", err)
		}

		router.Use(sentrygin.New(sentrygin.Options{
			Repanic: true,
		}))
	}

	router.POST("/", app.HandlePost)
	router.GET("/health", app.HandleHealth)
	router.GET("teams/:teamId/pins", app.HandleGetPins)
	router.GET("/teams/:teamId/channels", app.HandleGetChannels)
	router.GET("/teams/:teamId/channels/:channelId/pins", app.HandleGetPins)

	router.Run()
}
