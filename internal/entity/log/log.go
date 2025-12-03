package log

import (
	"fmt"
	"time"
)

type Log struct {
	IP        string    `log:"ip" parse:"string"`
	Identd    string    `log:"identd" parse:"string"`
	User      string    `log:"user" parse:"string"`
	Timestamp time.Time `log:"timestamp" parse:"apache_time"`
	Method    string    `log:"method" parse:"string"`
	Path      string    `log:"path" parse:"string"`
	Protocol  string    `log:"protocol" parse:"string"`
	Status    int       `log:"status" parse:"int"`
	Size      int       `log:"size" parse:"int"`
	Referer   string    `log:"referer" parse:"string"`
	UserAgent string    `log:"user_agent" parse:"string"`
}

func (l Log) String() string {
	return fmt.Sprintf(`	IP: %s
	Timestamp: %s
	Method: %s
	Path: %s
	Status: %d
	Size: %d
	Referer: %s
	UserAgent: %s`, l.IP, l.Timestamp, l.Method, l.Path, l.Status, l.Size, l.Referer, l.UserAgent)
}
