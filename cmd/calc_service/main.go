package main

import (
	"calc_service/internal/calculate"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result, err := calculate.Calc(req.Expression)
	var resp Response

	if err != nil {
		resp.Error = err.Error()
		w.WriteHeader(http.StatusUnprocessableEntity)
	} else {
		resp.Result = fmt.Sprintf("%f", result)
		w.WriteHeader(http.StatusOK)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/api/v1/calculate", calculateHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
