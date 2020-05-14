package model

//Logger type
type Logger struct {
	LevelLog  string `json:"level_log"`
	CreatedAt string `json:"created_at"`
	Message   string `json:"message"`
}
