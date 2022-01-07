package main

type user struct {
	ID   uint   `json: "id,omitempty"`
	Name string `json: "name,omitempty"`
}
