package main

type User struct {
	ID   uint   `json: "id,omitempty"`
	Name string `json: "name,omitempty"`
}
