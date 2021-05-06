package model

import "time"

type User struct {
	Id       uint64    `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Email    string    `json:"email,omitempty"`
	Password string    `json:"password,omitempty"`
	Token    string    `json:"token"`
	Created  time.Time `json:"created,omitempty"`
}

type ErrorMessage struct {
	Error   string `json:"error"`
	Message string `josn:"message"`
}
