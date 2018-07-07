package app

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/fortinj1354/Pin-Historian/models"
	"github.com/fortinj1354/Pin-Historian/settings"
	"github.com/gin-gonic/gin"
)

func HandlePost(c *gin.Context) {
	bytes, _ := c.GetRawData()
	requestTimestamp := c.GetHeader("X-Slack-Request-Timestamp")
	requestSignature := c.GetHeader("X-Slack-Signature")

	if validateSignature(string(bytes), requestTimestamp, settings.GetSlackSecret(), settings.GetSlackSigningVersion(), requestSignature) {
		var genericJson models.SlackGenericEventPost
		err := json.Unmarshal(bytes, &genericJson)
		if err == nil {
			if genericJson.Type == "url_verification" {
				HandleUrlVerification(c, bytes)
			} else if genericJson.Type == "event_callback" {
				c.JSON(http.StatusNoContent, nil)
				go HandleEventCallback(bytes)
			} else {
				c.JSON(http.StatusNoContent, nil)
			}
		} else {
			c.JSON(http.StatusNoContent, nil)
			panic(err)
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}
}

func validateSignature(body string, timeString string, secret string, version string, signature string) bool {
	match := false

	if body != "" && timeString != "" && signature != "" {
		timestamp, err := strconv.ParseFloat(timeString, 64)
		if err != nil {
			panic(err)
		}
		messageTime := time.Unix(int64(timestamp), 0)

		baseString := version + ":" + timeString + ":" + body
		mac := hmac.New(sha256.New, []byte(secret))
		mac.Write([]byte(baseString))
		calculated := mac.Sum(nil)
		byteSig, _ := hex.DecodeString(signature[3:])

		match = hmac.Equal(calculated, byteSig) && math.Abs(time.Since(messageTime).Seconds()) < 300
	}

	return match
}

func HandleGetPins(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	re := regexp.MustCompile(`^Bearer (.+)$`)
	apiKey := re.FindStringSubmatch(authorization)

	if apiKey == nil || apiKey[1] != settings.GetQueryAPIKey() {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	} else {
		teamId := c.Param("teamId")
		channelId := c.Param("channelId")
		userId := c.Query("userId")
		messageText := c.Query("messageText")
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

		if !startTime.IsZero() && !endTime.IsZero() && startTime.After(endTime) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "startTime cannot be after endTime"})
			return
		}

		response := QueryMessages(teamId, channelId, messageText, userId, startTime, endTime)
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

func HandleGetChannels(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	re := regexp.MustCompile(`^Bearer (.+)$`)
	apiKey := re.FindStringSubmatch(authorization)

	if apiKey == nil || apiKey[1] != settings.GetQueryAPIKey() {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	} else {
		teamId := c.Param("teamId")

		c.JSON(http.StatusOK, QueryChannels(teamId))
	}
}
