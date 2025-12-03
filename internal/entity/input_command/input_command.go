package inputcommand

import "time"

type InputCommand struct {
	Command  string
	FilePath string
	LogType  string

	IP        string
	Identd    string
	User      string
	Timestamp time.Time
	Method    string
	Path      string
	Protocol  string
	Status    int
	Size      int
	Referer   string
	UserAgent string
}
