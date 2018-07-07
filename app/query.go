package app

import (
	"sort"
	"time"

	"github.com/fortinj1354/Pin-Historian/models"
)

type PinResponse struct {
	TeamID       string           `json:"teamId"`
	ChannelID    string           `json:"channelId,omitempty"`
	StartTime    string           `json:"startTime,omitempty"`
	EndTime      string           `json:"endTime,omitempty"`
	Results      []models.Message `json:"results"`
	TotalResults int              `json:"totalResults"`
}

type ChannelResponse struct {
	Channels []models.Channel `json:"channels"`
}

func buildPinResponse(teamId string, channelId string, startTime time.Time, endTime time.Time, results []models.Message) *PinResponse {
	var response = &PinResponse{}
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

func QueryMessages(teamId string, channelId string, messageText string, userId string, startTime time.Time, endTime time.Time) *PinResponse {
	results := models.GetMessages(teamId, channelId, messageText, userId, startTime, endTime)

	return buildPinResponse(teamId, channelId, startTime, endTime, results)
}

func QueryChannels(teamId string) *ChannelResponse {
	results := models.GetChannels(teamId)

	response := ChannelResponse{
		Channels: results,
	}

	return &response
}
