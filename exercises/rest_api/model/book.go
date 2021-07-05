package model

//define your model
type Book struct {
	Id     string `json:"Id"`
	Title  string `json:"Title"`
	Author string `json:"Author"`
	Genre  string `json:"genre"`
	Isbn   string `json:"Isbn"`
}
