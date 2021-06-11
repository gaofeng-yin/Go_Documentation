package model

type Data struct {
	Metadata   Metadata `json:"metadata"`
	Pagination string   `json:"pagination"`
	Items      []Item   `json:"items"`
}

type Metadata struct {
	Page     uint `json:"page"`
	Per_page uint `json:"per_page"`
	Total    uint `json:"total"`
}

type Item struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}
