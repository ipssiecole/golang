package model

// Error is an error model.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
