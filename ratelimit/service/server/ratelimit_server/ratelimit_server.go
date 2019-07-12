package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/djinnes/riot/ratelimit/service/server"
)

var port = flag.Int("port", 8080, "server port")

func main() {
	flag.Parse()
	http.Handle("/", server.New())
	log.Println("listening on port", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
