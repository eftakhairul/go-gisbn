package gisbn

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"unsafe"
	"encoding/json"
)

//Base URL for calling  google book API
const baseURL = "https://www.googleapis.com"

//Tiny function handling for errr
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

var book Book

//GISBN object
type GISBN struct {
	isbn    string
	key     string
	country string	
}

//Set ISBN
//
//@param: isbn string
func (i *GISBN) SetISBN(isbn string) {
	i.isbn = isbn
}

//Core function for processing and fetching information 
//fomr Google Book API
func (i *GISBN) Fetch() {

	//formating API URL
	var requestUrl *url.URL
	requestUrl, err := url.Parse(baseURL)
	failOnError(err, "URL ")

	requestUrl.Path += "/books/v1/volumes"
	parameters := url.Values{}
	parameters.Add("q", fmt.Sprint("isbn:", i.isbn))
	parameters.Add("key", i.key)
	parameters.Add("country", i.country)
	requestUrl.RawQuery = parameters.Encode()

	response, err := http.Get(requestUrl.String())
	failOnError(err, "URL ")
	body, err := ioutil.ReadAll(response.Body)
	failOnError(err, "URL ")

	book, _ = getBook([]byte(body))	
}

//return key
func (i *GISBN)GetKey() string {
	return i.key
}

//return ISBN 13 time
func (i *GISBN)ISBN13() string {
	if len(book.IndustryIdentifiers) > 0 && book.IndustryIdentifiers[0].Type == "ISBN_13" {
		return book.IndustryIdentifiers[0].Identifier
	}

	if len(book.IndustryIdentifiers) > 0 && book.IndustryIdentifiers[1].Type == "ISBN_13" {
		return book.IndustryIdentifiers[1].Identifier
	}

	return ""
}

//return ISBN 10 time
func (i *GISBN)ISBN10() string {
	if len(book.IndustryIdentifiers) > 0 && book.IndustryIdentifiers[0].Type == "ISBN_10" {
		return book.IndustryIdentifiers[0].Identifier
	}

	if len(book.IndustryIdentifiers) > 0 && book.IndustryIdentifiers[1].Type == "ISBN_10" {
		return book.IndustryIdentifiers[1].Identifier
	}

	return ""
}

//return Thumbnail Image
func (i *GISBN)ThumbnailLink() string {	
	if unsafe.Sizeof(book) == 0 {
		return ""
	}

	return book.Images.Thumbnail
}

//return TotalPage
func (i *GISBN)TotalPage() int {	
	if unsafe.Sizeof(book) == 0 {
		return 0
	}

	return book.Page_count
}
//return Title
func (i *GISBN)Title() string {	
	if unsafe.Sizeof(book) == 0 {
		return ""
	}

	return book.Title
} 


//return all authors in array
func (i *GISBN)Publisher() string {	
	if unsafe.Sizeof(book) == 0 {
		return ""
	}

	return book.Publisher
} 



//Return book from from the main Json Response Object
//@param: body
//
//@return: Book
//@return: error
func getBook(body []byte) (Book, error) {
	var jsonresponse = new(JsonResponse)
	err := json.Unmarshal(body, &jsonresponse)
	if err != nil {
		fmt.Println("whoops:", err)
	}

	return jsonresponse.Items[0].Volumeinfo, err
}


