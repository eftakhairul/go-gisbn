package gisbn

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title       string   `json:"title"`
	Authors     []string `json:"authors"`
	Publisher   string   `json:"publisher"`
	Description string   `json:"description"`
	Page_count  int      `json:"pageCount"`
	Type        string   `json:"printType"`
	Categories  []string `json:"categories"`
}

type Item struct {
	Id         string `json:"id"`
	Selflink   string `json:"selfLink"`
	Volumeinfo Book   `json:"volumeInfo"`
}

type JsonResponse struct {
	Total int    `json:"totalItems"`
	Items []Item `json:"items"`
}

func getBook(body []byte) (Book, error) {
	var jsonresponse = new(JsonResponse)
	err := json.Unmarshal(body, &jsonresponse)
	if err != nil {
		fmt.Println("whoops:", err)
	}

	return jsonresponse.Items[0].Volumeinfo, err
}
