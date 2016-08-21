package main;

import(
      "github.com/eftakhairul/go-gisbn/gisbn"
      "fmt"
      )
        
func main() {
	g := gisbn.New("9780262033848", "AIzaSyDKepjfaVBRcgsnPALw5s2UNyfOk-1FHUU", "ca")
    g.Fetch()
    fmt.Println(g.Title())   //Introduction to Algorithms
}