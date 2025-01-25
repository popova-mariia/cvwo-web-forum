package models

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	PasswordHash string `json:"-"`
	Score        int    `json:"score"`
	IsAnonymous  bool   `json:"is_anonymous"`
}
