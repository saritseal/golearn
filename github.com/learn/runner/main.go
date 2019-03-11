package main

import (
	"flag"

	"algorithms"
)

var hostname string

func init() {
	const (
		defaultHostname = "localhost"
	)

	flag.StringVar(&hostname, "hostname", defaultHostname, "enter the hostname where you want to start the webserver")
}

func main() {
	algorithms.JaccardSimilarity()

}
