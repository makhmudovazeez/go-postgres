package models

type Stock struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Price   int    `json:"price"`
	Company string `json:"company"`
}
