package gisbn

//Image object
type ImageLinks struct {
	SmallThumbnail  string   `json:"smallThumbnail"`
	Thumbnail 		string   `json:"thumbnail"`
}

//ISBN Indentifier Object
type IndustryIdentifier struct {
	Type  			string   `json:"type"`
	Identifier 		string   `json:"identifier"`
}
//Book Object
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

//Item Object
//
//Main Json response holds this Item object
type Item struct {
	Id         string `json:"id"`
	Selflink   string `json:"selfLink"`
	Volumeinfo Book   `json:"volumeInfo"`
}

//Main Json Response
type JsonResponse struct {
	Total int    `json:"totalItems"`
	Items []Item `json:"items"`
}