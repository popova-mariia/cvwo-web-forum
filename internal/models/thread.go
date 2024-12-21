package models

type Thread struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	Topic           string `json:"topic"`
	CreatorID       int    `json:"creator_id"`
	ThemePreference string `json:"theme_preference"` // "light", "dark", or custom
}
