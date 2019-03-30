package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"server"
)

// var hostname string

var config server.ServerConfig

func init() {
	const (
		defaultHostname = "localhost"
	)

	config = server.ServerConfig{
		HostName: "localhost",
		Port:     8080,
	}

	// flag.StringVar(&hostname, "hostname", defaultHostname, "enter the hostname where you want to start the webserver")
}

func startTcpServer() {
	server.StartAccepting(server.StartTcp(config))
}

func startSendMessage() {
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		log.Fatal("error in connection")
	}

	for i := 0; i < 1000; i++ {

		_, err := bufio.NewWriter(conn).WriteString(fmt.Sprintf(""))

		if err != nil {
			log.Fatal("unable to write")
		}
	}
}

func main() {
	startTcpServer()
}
