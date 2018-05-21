package app

import (
	"net/http"
	"regexp"
	"time"

	"github.com/fortinj1354/Pin-Historian/models"
	"github.com/fortinj1354/Pin-Historian/settings"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func HandlePost(c *gin.Context) {
	var genericJson models.SlackGenericEventPost
	if err := c.ShouldBindBodyWith(&genericJson, binding.JSON); err == nil {
		if genericJson.Token != settings.GetSlackToken() {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		} else {
			if genericJson.Type == "url_verification" {
				HandleUrlVerification(c)
			} else if genericJson.Type == "event_callback" {
				HandleEventCallback(c)
			}
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func HandleGet(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	re := regexp.MustCompile(`^Bearer (.+)$`)
	apiKey := re.FindStringSubmatch(authorization)

	if apiKey == nil || apiKey[1] != settings.GetQueryAPIKey() {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	} else {
		teamId := c.Param("teamId")
		channelId := c.Param("channelId")
		var startTime time.Time
		var endTime time.Time

		if value, found := c.GetQuery("startTime"); found {
			var err error
			startTime, err = time.Parse(time.RFC3339, value)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}

		if value, found := c.GetQuery("endTime"); found {
			var err error
			endTime, err = time.Parse(time.RFC3339, value)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}

		results := QueryMessages(c, teamId, channelId, startTime, endTime)

		if results == nil {
			return
		}

		response := BuildResponse(teamId, channelId, startTime, endTime, *results)
		c.JSON(http.StatusOK, response)
	}
}

func HandleHealth(c *gin.Context) {
	err := models.Health()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
