package model

// @Description Payload for API errors
type Error struct {
	// Error code
	Code int `json:"Code"`
	// Error description
	Description string `json:"name" minLength:"4" maxLength:"16"`
}
