package biblio

import (
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
	"strings"
)

type Book struct {
	title  string
	url    string
	author string
	price  string
}

type PageInfo struct {
	page int
	sid  string
}

const URL = "www.biblio.com"

func Scrape(keyword string) {
	pageInfo := getIterationInfo(keyword)
	bookList := getPageBook(pageInfo.page, pageInfo.sid)
	fmt.Println(bookList)
}

func getIterationInfo(keyword string) PageInfo {
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
	return PageInfo{
		page: page,
		sid:  sid,
	}
}

func getPageBook(pageCount int, sid string) []Book {
	bookList := make([]Book, 1)

	for i := 1; i < pageCount; i++ {
		c := colly.NewCollector(colly.AllowedDomains(URL))
		c.OnHTML(".results.summary", func(element *colly.HTMLElement) {
			element.ForEach(".item", func(_ int, element *colly.HTMLElement) {
				title := element.ChildText(".basic-info .item-title .title")
				url := element.ChildAttr(".basic-info .item-title .title a", "href")
				author := element.ChildText(".basic-info .item-title .author")
				price := element.ChildText(".pricing .price-wrap .price span.item-price")

				bookList = append(bookList, Book{
					title:  title,
					url:    url,
					author: author,
					price:  price,
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
