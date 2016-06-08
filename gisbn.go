package gisbn

import (
	"fmt"
	 "net/http"
	 "net/url"
)

//Base URL for calling  google book API
const baseURL = "https://www.googleapis.com"


func failOnError(err error, msg string) {
  if err != nil {
    log.Fatalf("%s: %s", msg, err)
    panic(fmt.Sprintf("%s: %s", msg, err))
  }
}

// Book Object
//
// This contains a few private variables that are
// given accessor methods.
// This is the object that is created upon using
type Book struct {
  isbn_10         string
  isbn_13         string
  title           string
  authors         []string
  publishers      []string
  edition         string  
  cover           image.Image  
  published       time.Time
}




//ISBN object
//
type ISBN struct {
	isbn 	string 
	key 	string 
    country string   
}


//Set ISBN
func (i *isbn)setIsbn (isbn string) {
	i.isbn  = isbn
}


func (i *isbn) fetch(isbn string) {

	//formating API URL
	var requestUrl *url.URL
    requestUrl, err := url.Parse(baseURL)	
	failOnError(err, "URL ")

    requestUrl.Path += "/books/v1/volumes"
    parameters := url.Values{}
    parameters.Add("q", fmt.Sprint("isbn:", isbn))
    parameters.Add("key", i.key)
    parameters.Add("country", i.country)
    requestUrl.RawQuery = parameters.Encode()	

	response, err := http.Get(requestUrl.String())
}







