package main

import (
	"cheapest-bookiego/conventers"
	"cheapest-bookiego/crawlers/biblio"
	"cheapest-bookiego/models"
)

func main() {
	biblioCrawler := biblio.NewBiblio()

	crawlers := []models.ICrawler{biblioCrawler}
	for _, crawler := range crawlers {
		books := crawler.Scrape("python")
		conventers.ConvertToCsv(books, crawler.GetName())
	}
}
