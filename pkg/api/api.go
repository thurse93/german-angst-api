package api

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

type Opinion struct {
	Text string `json:"text"`
}

func generateOpinion(data map[string][]string) Opinion {
	adjective := data["adjective"][rand.Int()%len(data["adjective"])]
	subject := data["subject"][rand.Int()%len(data["subject"])]
	object := data["object"][rand.Int()%len(data["object"])]
	text := fmt.Sprintf("%s %s %s.", adjective, subject, object)
	return Opinion{text}
}

type OpinionHandler struct {
	Data map[string][]string
}

func (h *OpinionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		opinion := generateOpinion(h.Data)
		responseBody, err := json.Marshal(opinion)
		if err != nil {
			panic(err)
		}
		w.Write(responseBody)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
