package main

import (
	"github.com/fortinj1354/Pin-Historian/app"
	"github.com/fortinj1354/Pin-Historian/models"
	"github.com/fortinj1354/Pin-Historian/settings"
	"github.com/gin-gonic/gin"
)

func main() {
	settings.LoadSettings()
	router := gin.Default()
	models.MakeDB(settings.GetDatabaseName())

	router.POST("/", app.HandlePost)
	router.GET("/:teamId/:channelId", app.HandleGet)

	router.Run()
}
