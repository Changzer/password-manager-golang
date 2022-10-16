package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"password-manager/controllers"
)

func List(w http.ResponseWriter, r *http.Request) {
	todos, err := controllers.GetAll()
	if err != nil {
		log.Printf("deu erro %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
