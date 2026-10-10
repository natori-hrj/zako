package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	zako "github.com/natori-hrj/zako/go"
)

func main() {
	app := zako.New()

	app.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello, zako!")
	})

	server := &http.Server{
		Addr:              ":3000",
		Handler:           app,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
