package biblio

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"strconv"
	"strings"
)

import "cheapest-bookiego/models"

type CrawlerBiblio struct {
}

type PageInfo struct {
	page int
	sid  string
}

const URL = "www.biblio.com"

func NewBiblio() *CrawlerBiblio {
	return &CrawlerBiblio{}
}

func (crawler CrawlerBiblio) Scrape(keyword string) []models.Book {
	keyword = strings.Replace(keyword, " ", "%20", -1)
	pageInfo := crawler.getIterationInfo(keyword)
	bookList := crawler.getPageBook(pageInfo.page, pageInfo.sid)

	return bookList
}

func (crawler CrawlerBiblio) GetName() string {
	return "BIBLIO"
}

func (crawler CrawlerBiblio) getIterationInfo(keyword string) *PageInfo {
	var pages []string
	sid := ""

	c := colly.NewCollector(colly.AllowedDomains(URL))
	c.OnHTML("ul.pagination li", func(element *colly.HTMLElement) {
		element.ForEach("span", func(_ int, element *colly.HTMLElement) {
			if sid == "" {
				seperated := strings.Split(element.Attr("rel"), "sid=")
				sid = seperated[len(seperated)-1]
			}
			pages = append(pages, element.Text)
		})
	})
	customUrl := fmt.Sprintf("https://%s/search.php?stage=1&result_type=works&keyisbn=%s", URL, keyword)
	err := c.Visit(customUrl)
	if err != nil {
		fmt.Println("shit", err.Error())
	}

	page, _ := strconv.Atoi(pages[len(pages)-2])
	return &PageInfo{
		page: page,
		sid:  sid,
	}

}

func (crawler CrawlerBiblio) getPageBook(page int, sid string) []models.Book {
	bookList := make([]models.Book, 1)
	c := colly.NewCollector(colly.AllowedDomains(URL))
	if page > 5 {
		page = 5
	}

	for i := 1; i < page; i++ {
		c.OnHTML(".results.summary", func(element *colly.HTMLElement) {
			element.ForEach(".item", func(_ int, element *colly.HTMLElement) {
				title := element.ChildText(".basic-info .item-title .title")
				url := element.ChildAttr(".basic-info .item-title .title a", "href")
				author := element.ChildText(".basic-info .item-title .author")

				price := element.ChildText(".pricing .price-wrap .price span.item-price")
				price = strings.Split(price, "$")[1]
				log.Printf("[INFO]: Title: %s, Author: %s, Price: %s, Url: %s", title, author, price, url)

				bookList = append(bookList, models.Book{
					Title:  title,
					Url:    url,
					Author: author,
					Price:  price,
				})
			})
		})
		customUrl := fmt.Sprintf("https://%s/search.php?&page=%d&sid=%s", URL, i, sid)
		err := c.Visit(customUrl)
		if err != nil {
			fmt.Println("shit", err.Error())
		}
	}
	return bookList
}
