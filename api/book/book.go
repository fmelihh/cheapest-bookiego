package book

import (
	"cheapest-bookiego/crawler"
	"cheapest-bookiego/db/mongodb"
	"cheapest-bookiego/db/redis"
	"cheapest-bookiego/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func SearchBook(w http.ResponseWriter, r *http.Request) {
	keyword := r.FormValue("keyword")

	var response = models.BookJsonResponse{}
	mongoBookDatabase := mongodb.NewMongoBookModel()
	redisDatabase, _ := redis.NewRedisDatabase()

	if keyword == "" {
		response.Type = "error"
		response.Data = []models.BookDataResponse{}
		response.Message = "You are missing keyword parameter."

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := redisDatabase.Client.Get(keyword)
	if data.Val() == keyword {
		fmt.Print("bu daha önce aratılmış")
		mongoBooks, _ := mongoBookDatabase.FindByKeyword(keyword)

		// DB DEN ÇEKECEK KODU YAZ
	} else {
		redisDatabase.Client.Set(keyword, keyword, 0)

		bookData := crawler.Crawl(keyword)
		response = models.BookJsonResponse{
			Type:    "success",
			Data:    bookData,
			Message: "",
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}
