package book

import (
	"cheapest-bookiego/crawler"
	"cheapest-bookiego/db/redis"
	"cheapest-bookiego/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func SearchBook(w http.ResponseWriter, r *http.Request) {
	keyword := r.FormValue("keyword")

	var response = models.BookJsonResponse{}

	if keyword == "" {
		response = models.BookJsonResponse{
			Type:    "error",
			Data:    []models.BookDataResponse{},
			Message: "You are missing keyword parameter.",
		}
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		data, err := redis.GetKey(keyword)
		if err != nil {
			log.Printf("[ERR] redis err: %s", err.Error())
		}

		if data == keyword {
			fmt.Print("bu daha önce aratılmış")
			// DB DEN ÇEKECEK KODU YAZ
		} else {
			err = redis.SetKey(keyword)
			if err != nil {
				log.Printf("[ERR] redis err: %s", err.Error())
			}
			bookData := crawler.Crawl(keyword)
			response = models.BookJsonResponse{
				Type:    "success",
				Data:    bookData,
				Message: "",
			}
		}
		w.WriteHeader(http.StatusOK)
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatal("[ERROR]")
	}
}
