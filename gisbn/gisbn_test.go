package gisbn

import (
	"testing"
)

func TestFetch(t *testing.T) {
	g := GISBN{"0262033844", "AIzaSyDKepjfaVBRcgsnPALw5s2UNyfOk-1FHUU", "ca"}
	g.Fetch()
}
