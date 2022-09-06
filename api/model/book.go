package model

type Book struct {
	Id string `json:"id"`
	Title string `json:"title"`
	AuthorName string `json:"author_name"`
	Category string `json:"category"`
}
type CreateBook struct {
	Title  string `json:"title"`
	AuthorName string `json:"author_name"`
}

type UpdateBook struct {
	Title  string `json:"title"`
	AuthorName string `json:"author_name"`
}


type BookById struct {
	Id string `json:"id"`
}