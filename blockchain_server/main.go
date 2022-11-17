package main

import (
	"flag"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	port := flag.Uint("port", 5000, "TCP port number for Blockchain Server")
	flag.Parse()
	app := NewBlockchainServer(uint16(*port))
	app.Run()

}