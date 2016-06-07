package go-isbn

import (
	"fmt"
	 "net/http"
)

// Book Object
// This contains a few private variables that are
// given accessor methods.
// This is the object that is created upon using
// the method "GetBooks"
type Book struct {
  isbn            [2]string
  title           string
  authors         []string
  publishers      []string
  edition         string
  binding         string
  cover           image.Image
  msrp            float32
  published       time.Time
}