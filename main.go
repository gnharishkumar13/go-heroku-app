package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("This is a test heroku app deploy")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is up and running"))
	})
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}
