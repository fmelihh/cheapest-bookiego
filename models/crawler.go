package models

type ICrawler interface {
	Scrape(keyword string) ([]Book, error)
	GetName() string
}
