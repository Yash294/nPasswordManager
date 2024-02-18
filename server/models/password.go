package models

type Password struct {
	UserID     string `json:"userID"`
	PasswordID string `json:"passwordID"`
	Part       int    `json:"part"`
	Password   string `json:"password"`
}