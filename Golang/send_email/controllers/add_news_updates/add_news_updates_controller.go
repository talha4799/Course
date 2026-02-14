package add_news_updates_controller

import (
	"encoding/json"
	"login/db"
	add_news_update_db "login/db/add_newupdate_db"
	"net/http"
)

type AddNewsRequest struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	ServiceType string `json:"service_type"`
	Day         int    `json:"day"`
	Month       int    `json:"month"`
	Year        int    `json:"year"`
}

func AddNewsUpdates(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req AddNewsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate service type
	validService := false
	for _, s := range add_news_update_db.GetAllServiceTypes() {
		if s == req.ServiceType {
			validService = true
			break
		}
	}
	if !validService {
		http.Error(w, "Invalid service type", http.StatusBadRequest)
		return
	}

	update := &add_news_update_db.NewsUpdate{
		Title:       req.Title,
		Content:     req.Content,
		ServiceType: req.ServiceType,
		Day:         req.Day,
		Month:       req.Month,
		Year:        req.Year,
	}

	id, err := add_news_update_db.AddNewsUpdate(db.GetDB(), update)
	if err != nil {
		http.Error(w, "Failed to add news", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "News update added successfully",
		"id":      id,
	})
}
