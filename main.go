package main

import (
	_ "embed"
	"encoding/json"
	"log"
	"net/http"

	"github.com/thurse93/german-angst-api/pkg/api"
)

//go:embed data.json
var jsonData []byte

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s -> \n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	serverAddr := ":8080"

	data := make(map[string][]string)
	json.Unmarshal(jsonData, &data)

	server := &http.Server{
		Addr:    serverAddr,
		Handler: logRequest(http.DefaultServeMux),
	}

	http.Handle("/", &api.OpinionHandler{Data: data})

	log.Println("Listening on :8080")
	log.Fatal(server.ListenAndServe())

}
