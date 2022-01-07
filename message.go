package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	User    User   `json: "user,omitempty"`
	Content string `json: "content,omitempty"`
}

type Messages struct {
	messages []Message
}

func GetAllMessages(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all messages")
	json.NewEncoder(w).Encode(messages)
}
