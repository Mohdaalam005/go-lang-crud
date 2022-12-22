package models

type Book struct {
	ID          string  `json:"ID"`
	BookName    string  `json:"book_name"`
	BookPrice   float32 `json:"book_price"`
	Description string  `json:"description"`
	Author      *Author `json:"author"`
}

type Author struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	gender string `json:"gender"`
}
