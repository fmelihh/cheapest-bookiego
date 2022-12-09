package main

import (
	"cheapest-bookiego/conventers"
	"cheapest-bookiego/crawlers/bookdepository"
	"cheapest-bookiego/models"
)

func main() {
	//biblioCrawler := biblio.NewBiblio()
	bookDepositoryCrawler := bookdepository.NewCrawlerBookDepository()

	crawlers := []models.ICrawler{bookDepositoryCrawler}
	for _, crawler := range crawlers {
		books := crawler.Scrape("bbaax")
		conventers.ConvertToCsv(books, crawler.GetName())
	}
}
