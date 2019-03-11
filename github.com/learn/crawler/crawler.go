package crawler

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

// Crawl crawls a url passed
func Crawl(link string) *bytes.Buffer {

	resp, err := http.Get(link)

	if err != nil {
		log.Fatal("failed ro connect to the endpoint")
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal("failed ro connect to the endpoint", err)
	}

	log.Printf("Read data %s", link)
	return bytes.NewBuffer(data)
}
