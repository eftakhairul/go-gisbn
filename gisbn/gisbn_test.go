package gisbn

import (
	"testing"
)

func TestFetch(t *testing.T) {
	g := GISBN{"9780262033848", "AIzaSyDKepjfaVBRcgsnPALw5s2UNyfOk-1FHUU", "ca"}
	g.Fetch()
}
