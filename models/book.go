package models

type Book struct {
	Title  string
	Url    string
	Author string
	Price  string
}

type ICrawler interface {
	Scrape(keyword string) []Book
	GetName() string
}
