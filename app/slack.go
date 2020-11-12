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
	"github.com/parnurzeal/gorequest"
)

func HandleUrlVerification(c *gin.Context, body []byte) {
	var verificationJson models.SlackURLVerificationPost
	if err := json.Unmarshal(body, &verificationJson); err == nil {
		c.JSON(http.StatusOK, gin.H{"challenge": verificationJson.Challenge})
	} else {
		c.JSON(http.StatusNoContent, nil)
		panic(err)
	}
}

func HandleEventCallback(body []byte) {

	var eventCallbackJson models.SlackEventCallbackPost
	if err := json.Unmarshal(body, &eventCallbackJson); err == nil {
		eventType := eventCallbackJson.Event.Type
		if eventType == "pin_added" {
			handlePinnedItem(body)
		} else if eventType == "channel_rename" || eventType == "group_rename" {
			handleChannelRename(body)
		} else if eventType == "user_change" {
			handleUserChange(body)
		}
	} else {
		panic(err)
	}
}

func handlePinnedItem(body []byte) {
	var pinJson models.SlackPinPost
	if err := json.Unmarshal(body, &pinJson); err == nil {
		firstChar := pinJson.Event.Item.Channel[0]
		//Discard messages from DMs and private groups
		if firstChar == 'C' {
			timestamp, err := strconv.ParseFloat(pinJson.Event.Item.Message.Ts, 64)
			if err != nil {
				panic(err)
			}

			var groupName = resolveChannel(pinJson.TeamID, pinJson.Event.Item.Channel)

			message := &models.Message{
				EventID:     pinJson.EventID,
				TeamID:      pinJson.TeamID,
				ChannelID:   pinJson.Event.Item.Channel,
				ChannelName: groupName,
				UserID:      pinJson.Event.Item.Message.User,
				UserDisplay: resolveUser(pinJson.TeamID, pinJson.Event.Item.Message.User),
				MessageText: processMessageText(pinJson.TeamID, pinJson.Event.Item.Message.Text),
				MessageTime: time.Unix(int64(timestamp), 0)}

			models.SaveMessage(message)
		}
	} else {
		panic(err)
	}
}

func handleChannelRename(body []byte) {
	var channelJson models.SlackChannelRenamePost
	if err := json.Unmarshal(body, &channelJson); err == nil {
		channelModel := models.Channel{
			TeamID:      channelJson.TeamID,
			ChannelID:   channelJson.Event.Channel.ID,
			ChannelName: channelJson.Event.Channel.Name}

		models.UpdateChannel(&channelModel)
	} else {
		panic(err)
	}
}

func handleUserChange(body []byte) {
	var userJson models.SlackUserChangePost
	if err := json.Unmarshal(body, &userJson); err == nil {
		userModel := models.User{
			TeamID:      userJson.TeamID,
			UserID:      userJson.Event.User.ID,
			UserDisplay: userJson.Event.User.Profile.DisplayName}

		models.UpdateUser(&userModel)
	} else {
		panic(err)
	}
}

func resolveUser(teamId string, userId string) string {
	foundUser := models.GetUser(teamId, userId)

	if foundUser != nil {
		return foundUser.UserDisplay
	} else {
		var user models.SlackUserRequest

		request := gorequest.New()
		_, _, err := request.Get("https://slack.com/api/users.profile.get").
			Set("Authorization", "Bearer "+settings.GetSlackOAuth()).
			Query("user=" + userId).
			EndStruct(&user)
		if err != nil {
			panic(err)
		}

		userModel := models.User{
			TeamID: teamId,
			UserID: userId}

		if user.Profile.BotID == "" {
			userModel.UserDisplay = user.Profile.DisplayName
		} else {
			userModel.UserDisplay = user.Profile.RealName
		}

		models.SaveUser(&userModel)

		return userModel.UserDisplay
	}
}

func resolveChannel(teamId string, channelId string) string {
	foundChannel := models.GetChannel(teamId, channelId)

	if foundChannel != nil {
		return foundChannel.ChannelName
	} else {
		var channel models.SlackChannelRequest

		request := gorequest.New()
		_, _, err := request.Get("https://slack.com/api/conversations.info").
			Set("Authorization", "Bearer "+settings.GetSlackOAuth()).
			Query("channel=" + channelId).
			EndStruct(&channel)
		if err != nil {
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
	re := regexp.MustCompile(`<@(U.{8})>`)
	res := re.FindAllStringSubmatch(messageText, -1)

	for _, match := range res {
		userDisplay := resolveUser(teamId, match[1])
		tempRe := regexp.MustCompile(match[0])
		messageText = tempRe.ReplaceAllString(messageText, "@"+userDisplay)
	}

	return messageText
}
