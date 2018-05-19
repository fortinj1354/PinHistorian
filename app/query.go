package app

import (
	"github.com/fortinj1354/Pin-Historian/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"time"
)

type Response struct {
	TeamID       string           `json:"teamId"`
	ChannelID    string           `json:"channelId"`
	StartTime    string           `json:"startTime"`
	EndTime      string           `json:"endTime"`
	Results      []models.Message `json:"results"`
	TotalResults int              `json:"totalResults"`
}

func BuildResponse(teamId string, channelId string, startTime time.Time, endTime time.Time, results []models.Message) *Response {
	var response = &Response{}
	response.TeamID = teamId
	response.ChannelID = channelId

	if !startTime.IsZero() {
		response.StartTime = startTime.Format(time.RFC3339)
	}

	if !endTime.IsZero() {
		response.EndTime = endTime.Format(time.RFC3339)
	}

	if len(results) > 1 {
		sort.Slice(results, func(i int, j int) bool { return results[i].MessageTime.Before(results[j].MessageTime) })
	}

	if len(results) < 1 {
		response.Results = []models.Message{}
		response.TotalResults = 0
	} else {
		response.Results = results
		response.TotalResults = len(results)
	}

	return response
}

func QueryMessages(c *gin.Context, teamId string, channelId string, startTime time.Time, endTime time.Time) *[]models.Message {
	var results []models.Message

	if startTime.IsZero() && endTime.IsZero() {
		results = models.GetAllMessages(teamId, channelId)
	} else {
		if !startTime.IsZero() {
			if !endTime.IsZero() {
				if startTime.After(endTime) {
					c.JSON(http.StatusBadRequest, gin.H{"error": "startTime cannot be after endTime"})
					return nil
				} else {
					results = models.GetMessagesInRange(teamId, channelId, startTime, endTime)
				}
			} else {
				results = models.GetMessagesStartTime(teamId, channelId, startTime)
			}
		} else {
			results = models.GetMessagesEndTime(teamId, channelId, endTime)
		}
	}

	return &results
}
