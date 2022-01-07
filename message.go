package main

import "time"

type message struct {
	User      User      `json: "user,omitempty"`
	Content   string    `json: "content,omitempty"`
	Timestamp time.Time `json: "timestamp,omitempty"`
}
