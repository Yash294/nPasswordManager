package models

type Password struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Part 	 int	`json:"part"`
	Password []byte `json:"password"`
}