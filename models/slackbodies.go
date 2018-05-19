package models

type GenericSlackPost struct {
	Token string `json:"token" binding:"required,min=1"`
	Type  string `json:"type" binding:"required,eq=url_verification|eq=event_callback"`
}

type UrlVerificationSlackPost struct {
	Challenge string `json:"challenge" binding:"required,min=1"`
}

type PinSlackPost struct {
	TeamID string `json:"team_id" binding:"required,min=1"`
	Event  struct {
		Type string `json:"type" binding:"required,eq=pin_added"`
		Item struct {
			Type    string `json:"type" binding:"required,eq=message"`
			Channel string `json:"channel" binding:"required,min=1"`
			Message struct {
				User string `json:"user" binding:"required,min=1"`
				Text string `json:"text" binding:"required,min=1"`
				Ts   string `json:"ts" binding:"required,min=1"`
			} `json:"message" binding:"required"`
		} `json:"item" binding:"required"`
	} `json:"event" binding:"required"`
	EventID string `json:"event_id" binding:"required,min=1"`
}

type SlackUserRequest struct {
	Profile struct {
		DisplayName string `json:"display_name"`
	} `json:"profile"`
}

type SlackChannelRequest struct {
	Channel struct {
		Name string `json:"name"`
	} `json:"channel"`
}
