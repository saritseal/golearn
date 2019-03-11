package crawler

import (
	"log"
	"strconv"
	"testing"
)

func TestCrawl(t *testing.T) {
	Crawl("http://www.cnn.com")
}

func BenchmarkCrawl(t *testing.B) {

	var testData = []struct {
		name string
		url  string
	}{
		{"cnn", "http://www.cnn.com"},
		{"bbc", "http://www.bbc.com"},
		{"ndtv", "http://www.ndtv.com"},
	}

	for n := 0; n < 10; n++ {
		for _, tt := range testData {
			buf := Crawl(tt.url)
			log.Printf(strconv.Itoa(len(string(buf.Bytes()))))
		}
	}
}
