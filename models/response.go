package models

type BookDataResponse struct {
	BookData []Book `json:"book_data"`
	Message  string `json:"message"`
}

type BookJsonResponse struct {
	Type    string             `json:"type"`
	Data    []BookDataResponse `json:"data"`
	Message string             `json:"message"`
}
