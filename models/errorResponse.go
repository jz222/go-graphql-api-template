package models

type Error struct {
	Code    int    `json:"code"`
	Message string `message:"message"`
}
