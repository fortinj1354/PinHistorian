package models

type SlackGenericEventPost struct {
	Token string `json:"token" binding:"required,min=1"`
	Type  string `json:"type" binding:"required,eq=url_verification|eq=event_callback"`
}

type SlackURLVerificationPost struct {
	Challenge string `json:"challenge" binding:"required,min=1"`
}

type SlackEventCallbackPost struct {
	Event struct {
		Type string `json:"type" binding:"required,eq=pin_added|eq=channel_rename|eq=user_change"`
	} `json:"event" binding:"required"`
}

type SlackPinPost struct {
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

type SlackChannelRenamePost struct {
	TeamID string `json:"team_id" binding:"required,min=1"`
	Event  struct {
		Type    string `json:"type" binding:"required,eq=channel_rename"`
		Channel struct {
			ID   string `json:"id" binding:"required"`
			Name string `json:"name" binding:"required"`
		} `json:"channel" binding:"required"`
	} `json:"event" binding:"required"`
}

type SlackUserChangePost struct {
	TeamID string `json:"team_id" binding:"required,min=1"`
	Event  struct {
		Type string `json:"type" binding:"required,eq=user_change"`
		User struct {
			ID      string `json:"id" binding:"required"`
			Profile struct {
				DisplayName string `json:"display_name" binding:"required"`
			} `json:"profile" binding:"required"`
		} `json:"user" binding:"required"`
	} `json:"event" binding:"required"`
}

type SlackUserRequest struct {
	Profile struct {
		DisplayName string `json:"display_name"`
		BotID       string `json:"bot_id"`
		RealName    string `json:"real_name"`
	} `json:"profile"`
}

type SlackChannelRequest struct {
	Channel struct {
		Name string `json:"name"`
	} `json:"channel"`
}
