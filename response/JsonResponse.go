package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type JsonResponse struct{}

func (jr JsonResponse) Send(w http.ResponseWriter, data interface{}, statusCode int) {
	body, err := json.Marshal(data)
	if err != nil {
		log.Printf("Could not encode JSON data: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Accept", "application/json; charset=UTF-8")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	_, err = w.Write(body)
	if err != nil {
		log.Printf("Could not process the request: %v", err)
		return
	}
}
