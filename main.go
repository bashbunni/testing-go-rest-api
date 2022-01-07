package main

import (
	"fmt"
	"log"
	"net/http"
)

var messages []Message

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to the message board")
	fmt.Println("Endpoint access success")
}

func handleRequests() {
	http.HandleFunc("/", welcome)
	http.HandleFunc("/messages", GetAllMessages)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	messages = []Message{
		{User{1, "bob"}, "Hello!"},
		{User{2, "tom"}, "Yo"},
		{User{3, "joe"}, "idk"},
	}
	handleRequests()
}
