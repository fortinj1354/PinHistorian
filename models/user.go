package models

import "database/sql"

type User struct {
	TeamID      string
	UserID      string
	UserDisplay string
}

func SaveUser(user *User) {
	insertStmt, err := db.Prepare("INSERT INTO users(teamID, userID, userDisplay) values (?,?,?)")
	if err != nil {
		panic(err)
	}
	defer insertStmt.Close()
	insertStmt.Exec(user.TeamID, user.UserID, user.UserDisplay)
}

func GetUser(teamId string, userId string) *User {
	selectStmt, err := db.Prepare("SELECT teamID, userID, userDisplay FROM users WHERE teamID = ? AND userID = ?")
	if err != nil {
		panic(err)
	}
	defer selectStmt.Close()
	user := &User{}
	err = selectStmt.QueryRow(teamId, userId).Scan(&user.TeamID, &user.UserID, &user.UserDisplay)
	switch {
	case err == sql.ErrNoRows:
		return nil
	case err != nil:
		panic(err)
	default:
	}
	return user
}

func UpdateUser(user *User) {
	updateChannelStmt, err := db.Prepare("UPDATE users SET userDisplay = ? WHERE teamID = ? AND userID = ?")
	if err != nil {
		panic(err)
	}
	defer updateChannelStmt.Close()
	updateChannelStmt.Exec(user.UserDisplay, user.TeamID, user.UserID)

	updateMessagesStmt, err := db.Prepare("UPDATE messages SET userDisplay = ? WHERE teamID = ? AND userID = ?")
	if err != nil {
		panic(err)
	}
	defer updateMessagesStmt.Close()
	updateMessagesStmt.Exec(user.UserDisplay, user.TeamID, user.UserID)
}
