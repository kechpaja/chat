package main

import (
	"fmt"
)

type UserID string

type Message struct {
	Msg string
	User UserID
	// TODO Uid int64?
}

// Gets a string from the user and returns it
func GetUserInput() (string, bool) {
	var s string
	i, err := fmt.Scan(&s)
	if err != nil {
		return nil, false
	}
	return s, true
}

func CreateMessage(msg string, user UserID) Message {
	m := new(Message)
	m.Msg = msg
	m.User = user
	return m
}

func DisplayToUser(msg Message) {
	fmt.Print("%s: %s\n", msg.User, msg.Msg) // TODO time? 
}


