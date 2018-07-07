package models

import (
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
)

type Message struct {
	EventID     string    `json:"eventId"`
	TeamID      string    `json:"teamId"`
	ChannelID   string    `json:"channelId"`
	ChannelName string    `json:"channelName"`
	UserID      string    `json:"userId"`
	UserDisplay string    `json:"userDisplay"`
	MessageText string    `json:"messageText"`
	MessageTime time.Time `json:"messageTime"`
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

func GetMessages(teamId string, channelId string, messageText string, userId string, startTime time.Time, endTime time.Time) []Message {

	//Generate SQL
	selectSql := sq.Select("eventID, teamID, channelID, channelName, userID, userDisplay, messageText, messageTime").From("messages").Where(sq.Eq{"teamId": teamId})

	if channelId != "" {
		selectSql = selectSql.Where(sq.Eq{"channelID": channelId})
	}

	if messageText != "" {
		selectSql = selectSql.Where("messageText LIKE ?", fmt.Sprint("%", messageText, "%"))
	}

	if userId != "" {
		selectSql = selectSql.Where(sq.Eq{"userId": userId})
	}

	if !startTime.IsZero() {
		selectSql = selectSql.Where(sq.GtOrEq{"messageTime": startTime})
	}

	if !endTime.IsZero() {
		selectSql = selectSql.Where(sq.LtOrEq{"messageTime": endTime})
	}

	genSql, args, err := selectSql.ToSql()

	//Use SQL
	selectStmt, err := db.Prepare(genSql)
	if err != nil {
		panic(err)
	}

	defer selectStmt.Close()
	rows, selectErr := selectStmt.Query(args...)
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
