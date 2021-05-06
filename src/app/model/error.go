package model

type Error struct {
	Error   string `json:"error"`
	Message string `joson:"message"`
}
