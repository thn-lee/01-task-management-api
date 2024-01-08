package models

import (
	"log"
)

// Error represents an error that occurred while handling a request.
type Error struct {
	Code    int         `json:"code"`
	Source  interface{} `json:"source,omitempty"`
	Title   string      `json:"title,omitempty"`
	Message string      `json:"message,omitempty"`
}

func (e *Error) Error() (errStr string) {
	return e.Message
}

func (e *Error) Log() {
	log.Printf("source: %+s \nerr: %+s \n", e.Source, e.Message)
}
