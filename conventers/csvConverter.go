package conventers

import (
	"cheapest-bookiego/models"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

func ConvertToCsv(records []models.Book, webName string) {
	currentTime := time.Now().Format("02-01-2006")
	filename := fmt.Sprintf("%s-%s.csv", webName, currentTime)

	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		log.Fatal("failed to open file", err)
	}

	w := csv.NewWriter(file)
	defer w.Flush()

	var data [][]string
	data = append(data, []string{"AUTHOR", "PRICE", "TITLE", "URL"})
	for _, book := range records {
		row := []string{book.Author, book.Price, book.Title, book.Url}
		data = append(data, row)
	}

	_ = w.WriteAll(data)

}
