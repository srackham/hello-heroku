package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello Heroku!\n")
	}

	http.HandleFunc("/", helloHandler)
	print("listening on port", port, "\n")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
