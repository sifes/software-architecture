package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Response struct {
  CurrentTime string `json:"current_time"`
}
func handleTimeRequest(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodGet {
    http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
    return
  }

  timeData := Response{CurrentTime: time.Now().Format(time.RFC3339)}

  w.Header().Set("Content-Type", "application/json")

  if err := json.NewEncoder(w).Encode(timeData); err != nil {
    http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
    log.Printf("Encoding error: %v", err)
  }
}
