package model

type Point struct {
	Id         int    `json: "id"`
	Created_at string `json: "created_at"`
	Updated_at string `json: "updated_at"`
	Active     bool   `json: "active"`
}
