package main

import (
	"fmt"
	"log"
	"net/http"

	"./client"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/read_article/", articleHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1><center>grpc article client</center></h1>")
	fmt.Fprintf(w, "run http://localhost:8080/read_article?url={URL} to parse article")
}

func articleHandler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["url"]

	if !ok || len(keys[0]) < 1 {
		fmt.Fprintf(w, "Url Param 'url' is missing")
		return
	}

	url := keys[0]
	var clientResponse = client.SendRequest(url)
	fmt.Fprintf(w, clientResponse)
	return

}
