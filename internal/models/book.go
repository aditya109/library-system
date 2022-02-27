package models

type Book struct {
	Id          string  `json:"_id,omitempty"`
	BookId      int     `json:"bookid,omitempty"`
	BookName    string  `json:"bookname,omitempty"`
	ISBN        string  `json:"isbn,omitempty"`
	BookAuthor  string  `json:"bookauthor,omitempty"`
	Price       float64 `json:"price,omitempty"`
	IsAvailable bool    `json:"isAvailable,omitempty"`
}
