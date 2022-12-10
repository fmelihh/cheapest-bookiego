package crawler

import (
	"cheapest-bookiego/conventers"
	"cheapest-bookiego/crawler/biblio"
	"cheapest-bookiego/crawler/bookdepository"
	"cheapest-bookiego/models"
	"fmt"
)

func Crawl(keyword string) []models.BookDataResponse {
	biblioCrawler := biblio.NewBiblio()
	bookDepositoryCrawler := bookdepository.NewCrawlerBookDepository()

	var bookResponse []models.BookDataResponse

	crawlers := []models.ICrawler{biblioCrawler, bookDepositoryCrawler}
	for _, crawler := range crawlers {

		books, err := crawler.Scrape(keyword)
		if err != nil {
			errMsg := fmt.Sprintf("%s: %s", crawler.GetName(), "Data Not Found (!)")
			bookResponse = append(bookResponse, models.BookDataResponse{
				BookData: books,
				Message:  errMsg,
			})
			fmt.Printf("[ERR]: msg = %s", err.Error())
		} else {
			successMessage := fmt.Sprintf("%s: %s, [%d]", crawler.GetName(), "Data Count", len(books))
			bookResponse = append(bookResponse, models.BookDataResponse{
				BookData: books,
				Message:  successMessage,
			})
		}
		conventers.ConvertToCsv(books, crawler.GetName())
	}
	return bookResponse
}
