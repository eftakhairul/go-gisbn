package gisbn

import (
	"encoding/json"
	"fmt"
)

type ImageLinks struct {
	SmallThumbnail  string   `json:"smallThumbnail"`
	Thumbnail 		string   `json:"thumbnail"`
}

type IndustryIdentifier struct {
	Type  			string   `json:"type"`
	Identifier 		string   `json:"identifier"`
}

type Book struct {
	Title       		string   				`json:"title"`
	Authors     		[]string 				`json:"authors"`
	Publisher   		string   				`json:"publisher"`
	Description 		string   				`json:"description"`
	Page_count  		int      				`json:"pageCount"`
	Type        		string   				`json:"printType"`
	Categories  		[]string 				`json:"categories"`
	Images				ImageLinks  			`json:"imageLinks"`
	IndustryIdentifiers []IndustryIdentifier	`json:"industryIdentifiers"`		
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


func (b *Book)ISBN_13() string {
	if len(b.IndustryIdentifiers) > 0 && b.IndustryIdentifiers[0].Type == "ISBN_13" {
		return b.IndustryIdentifiers[0].Identifier
	}

	if len(b.IndustryIdentifiers) > 0 && b.IndustryIdentifiers[1].Type == "ISBN_13" {
		return b.IndustryIdentifiers[1].Identifier
	}

	return "Sorry  not found"
}


func (b *Book)ISBN_10() string {
	if len(b.IndustryIdentifiers) > 0 && b.IndustryIdentifiers[0].Type == "ISBN_10" {
		return b.IndustryIdentifiers[0].Identifier
	}

	if len(b.IndustryIdentifiers) > 0 && b.IndustryIdentifiers[1].Type == "ISBN_10" {
		return b.IndustryIdentifiers[1].Identifier
	}

	return "Sorry  not found"
}

func (b *Book)ThumbnailLink() string {
	
	return b.Images.Thumbnail
}