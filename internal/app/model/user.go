package model

type Users struct {
	Id         int    `json: "id"`
	Created_at string `json: "create_at"`
	Updated_at string `json: "updated_at"`
	Login      string `json: "login"`
	Apikey     string `json: "apikey"`
}
