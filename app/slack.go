package app

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/fortinj1354/Pin-Historian/models"
	"github.com/fortinj1354/Pin-Historian/settings"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/parnurzeal/gorequest"
)

func HandleUrlVerification(c *gin.Context) {
	var verificationJson models.UrlVerificationSlackPost
	if err := c.ShouldBindBodyWith(&verificationJson, binding.JSON); err == nil {
		c.JSON(http.StatusOK, gin.H{"challenge": verificationJson.Challenge})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func HandlePinnedItem(c *gin.Context) {
	var pinJson models.PinSlackPost
	if err := c.ShouldBindBodyWith(&pinJson, binding.JSON); err == nil {
		timestamp, err := strconv.ParseFloat(pinJson.Event.Item.Message.Ts, 64)
		if err != nil {
			panic(err)
		}

		message := &models.Message{
			EventID:     pinJson.EventID,
			TeamID:      pinJson.TeamID,
			ChannelID:   pinJson.Event.Item.Channel,
			ChannelName: resolveChannel(pinJson.TeamID, pinJson.Event.Item.Channel),
			UserID:      pinJson.Event.Item.Message.User,
			UserDisplay: resolveUser(pinJson.TeamID, pinJson.Event.Item.Message.User),
			MessageText: processMessageText(pinJson.TeamID, pinJson.Event.Item.Message.Text),
			MessageTime: time.Unix(int64(timestamp), 0)}

		models.SaveMessage(message)

		c.JSON(http.StatusNoContent, nil)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func resolveUser(teamId string, userId string) string {
	foundUser := models.GetUser(teamId, userId)

	if foundUser != nil {
		return foundUser.UserDisplay
	} else {
		request := gorequest.New()
		resp, _, err := request.Get("https://slack.com/api/users.profile.get").
			Set("Authorization", "Bearer "+settings.GetSlackOAuth()).
			Query("user=" + userId).
			End()
		if err != nil {
			panic(err)
		}

		var user models.SlackUserRequest
		jerr := json.NewDecoder(resp.Body).Decode(&user)
		if jerr != nil {
			panic(err)
		}

		userModel := models.User{
			TeamID:      teamId,
			UserID:      userId,
			UserDisplay: user.Profile.DisplayName}

		models.SaveUser(&userModel)

		return user.Profile.DisplayName
	}
}

func resolveChannel(teamId string, channelId string) string {
	foundChannel := models.GetChannel(teamId, channelId)

	if foundChannel != nil {
		return foundChannel.ChannelName
	} else {
		request := gorequest.New()
		resp, _, err := request.Get("https://slack.com/api/channels.info").
			Set("Authorization", "Bearer "+settings.GetSlackOAuth()).
			Query("channel=" + channelId).
			End()
		if err != nil {
			panic(err)
		}

		var channel models.SlackChannelRequest
		jerr := json.NewDecoder(resp.Body).Decode(&channel)
		if jerr != nil {
			panic(err)
		}

		channelModel := models.Channel{
			TeamID:      teamId,
			ChannelID:   channelId,
			ChannelName: channel.Channel.Name}

		models.SaveChannel(&channelModel)

		return channel.Channel.Name
	}
}

func processMessageText(teamId string, messageText string) string {
	re, err := regexp.Compile(`<@(U.{8})>`)
	if err != nil {
		panic(err)
	}

	res := re.FindAllStringSubmatch(messageText, -1)

	for _, match := range res {
		userDisplay := resolveUser(teamId, match[1])
		tempRe, err := regexp.Compile(match[0])
		if err != nil {
			panic(err)
		}

		messageText = tempRe.ReplaceAllString(messageText, "@"+userDisplay)
	}

	return messageText
}
