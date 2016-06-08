package gisbn

import (
	"fmt"
	"log"
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





//ISBN object
//
type ISBN struct {
	isbn 	string 
	key 	string 
    country string   
}


//Set ISBN
func (i *ISBN)setIsbn (isbn string) {
	i.isbn  = isbn
}


func (i *ISBN) fetch() {

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
	fmt.Println(response)
}







