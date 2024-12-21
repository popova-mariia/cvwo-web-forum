package models

type User struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Score       int    `json:"score"`
	IsAnonymous bool   `json:"is_anonymous"`
}
