package bookdepository

import (
	"cheapest-bookiego/models"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"strings"
)

type CrawlerBookDepository struct {
}

const URL = "www.bookdepository.com"

func NewCrawlerBookDepository() *CrawlerBookDepository {
	return &CrawlerBookDepository{}
}

func (crawler CrawlerBookDepository) GetName() string {
	return "BOOKDEPOSITORY"
}

func (crawler CrawlerBookDepository) Scrape(keyword string) []models.Book {
	keyword = strings.Replace(keyword, " ", "%20", -1)
	bookList := crawler.getPageBook(keyword)
	return bookList
}

func (crawler CrawlerBookDepository) getPageBook(keyword string) []models.Book {
	bookList := make([]models.Book, 0)
	counter := 0
	c := colly.NewCollector(colly.AllowedDomains(URL))
	lastLength := 0

	for lastLength == len(bookList) {
		c.OnHTML(".tab.search", func(element *colly.HTMLElement) {
			element.ForEach(".book-item", func(_ int, element *colly.HTMLElement) {
				title := element.ChildText(".item-info h3.title a")
				url := "https://" + URL + element.ChildAttr(".item-info h3.title a", "href")
				author := element.ChildText(".item-info p.author")
				price := element.ChildText(".item-info .price-wrap span.sale-price")

				log.Printf("[INFO]: Title: %s, Author: %s, Price: %s, Url: %s", title, author, price, url)

				bookList = append(bookList, models.Book{
					Title:  title,
					Url:    url,
					Author: author,
					Price:  price,
				})
			})
			counter = counter + 1
			lastLength = len(bookList)
		})
		customUrl := fmt.Sprintf("https://%s/search?searchTerm=%s&page=%d", URL, keyword, counter)
		err := c.Visit(customUrl)
		if err != nil {
			fmt.Println("shit", err.Error())
			break
		}
	}
	return bookList
}
