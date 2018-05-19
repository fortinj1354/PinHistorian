package models

import (
	"database/sql"
	"time"
)

type Message struct {
	EventID     string
	TeamID      string
	ChannelID   string
	ChannelName string
	UserID      string
	UserDisplay string
	MessageText string
	MessageTime time.Time
}

func SaveMessage(message *Message) {
	selectStmt, err := db.Prepare("SELECT count(*) FROM messages WHERE teamID = ? AND channelID = ? AND messageTime = ?")
	if err != nil {
		panic(err)
	}

	defer selectStmt.Close()
	var count int
	selectStmt.QueryRow(message.TeamID, message.ChannelID, message.MessageTime).Scan(&count)

	if count == 0 {
		insertStmt, err := db.Prepare("INSERT INTO messages(eventID, teamID, channelID, channelName, userID, userDisplay, messageText, messageTime) VALUES (?,?,?,?,?,?,?,?)")
		if err != nil {
			panic(err)
		}
		defer insertStmt.Close()
		insertStmt.Exec(message.EventID, message.TeamID, message.ChannelID, message.ChannelName, message.UserID, message.UserDisplay, message.MessageText, message.MessageTime)
	}
}

func GetAllMessages(teamId string, channelId string) []Message {
	selectStmt, err := db.Prepare("SELECT eventID, teamID, channelID, channelName, userID, userDisplay, messageText, messageTime FROM messages WHERE teamID = ? AND channelID = ?")
	if err != nil {
		panic(err)
	}

	defer selectStmt.Close()
	rows, selectErr := selectStmt.Query(teamId, channelId)
	var messages []Message

	switch {
	case selectErr == sql.ErrNoRows:
		return messages
	case selectErr != nil:
		panic(err)
	default:
	}

	for rows.Next() {
		message := Message{}
		rows.Scan(&message.EventID, &message.TeamID, &message.ChannelID, &message.ChannelName, &message.UserID, &message.UserDisplay, &message.MessageText, &message.MessageTime)
		messages = append(messages, message)
	}

	return messages
}

func GetMessagesStartTime(teamId string, channelId string, startTime time.Time) []Message {
	selectStmt, err := db.Prepare("SELECT eventID, teamID, channelID, channelName, userID, userDisplay, messageText, messageTime FROM messages WHERE teamID = ? AND channelID = ? AND messageTime > ?")
	if err != nil {
		panic(err)
	}

	defer selectStmt.Close()
	rows, selectErr := selectStmt.Query(teamId, channelId, startTime)
	var messages []Message

	switch {
	case selectErr == sql.ErrNoRows:
		return messages
	case selectErr != nil:
		panic(err)
	default:
	}

	for rows.Next() {
		message := Message{}
		rows.Scan(&message.EventID, &message.TeamID, &message.ChannelID, &message.ChannelName, &message.UserID, &message.UserDisplay, &message.MessageText, &message.MessageTime)
		messages = append(messages, message)
	}

	return messages
}

func GetMessagesEndTime(teamId string, channelId string, endTime time.Time) []Message {
	selectStmt, err := db.Prepare("SELECT eventID, teamID, channelID, channelName, userID, userDisplay, messageText, messageTime FROM messages WHERE teamID = ? AND channelID = ? AND messageTime < ?")
	if err != nil {
		panic(err)
	}

	defer selectStmt.Close()
	rows, selectErr := selectStmt.Query(teamId, channelId, endTime)
	var messages []Message

	switch {
	case selectErr == sql.ErrNoRows:
		return messages
	case selectErr != nil:
		panic(err)
	default:
	}

	for rows.Next() {
		message := Message{}
		rows.Scan(&message.EventID, &message.TeamID, &message.ChannelID, &message.ChannelName, &message.UserID, &message.UserDisplay, &message.MessageText, &message.MessageTime)
		messages = append(messages, message)
	}

	return messages
}

func GetMessagesInRange(teamId string, channelId string, startTime time.Time, endTime time.Time) []Message {
	selectStmt, err := db.Prepare("SELECT eventID, teamID, channelID, channelName, userID, userDisplay, messageText, messageTime FROM messages WHERE teamID = ? AND channelID = ? AND messageTime > ? AND messageTime < ?")
	if err != nil {
		panic(err)
	}

	defer selectStmt.Close()
	rows, selectErr := selectStmt.Query(teamId, channelId, startTime, endTime)
	var messages []Message

	switch {
	case selectErr == sql.ErrNoRows:
		return messages
	case selectErr != nil:
		panic(err)
	default:
	}

	for rows.Next() {
		message := Message{}
		rows.Scan(&message.EventID, &message.TeamID, &message.ChannelID, &message.ChannelName, &message.UserID, &message.UserDisplay, &message.MessageText, &message.MessageTime)
		messages = append(messages, message)
	}

	return messages
}
