package models

type Role struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"` //eg  "Frontend Developer", "Backend Developer", "UI/UX Designer", "Product Manager", "IDK lmao"
	Leader bool   `json:"leader"`
}
