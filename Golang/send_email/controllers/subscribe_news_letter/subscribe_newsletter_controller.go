package subscribe_news_letter_controller

import (
	"encoding/json"
	"login/db"
	add_news_update_db "login/db/add_newupdate_db"
	"login/db/subscription_db"
	"net/http"
)

type SubscribeRequest struct {
	Email      string `json:"email"`
	FullName   string `json:"full_name"`
	ServiceOpt string `json:"service_opt"`
}

type UnsubscribeRequest struct {
	Email      string `json:"email"`
	ServiceOpt string `json:"service_opt"`
}

func SubscribeNewsletter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handleSubscribe(w, r)
	case http.MethodDelete:
		handleUnsubscribe(w, r)
	case http.MethodGet:
		handleGetServices(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleSubscribe(w http.ResponseWriter, r *http.Request) {
	var req SubscribeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate service option
	valid := false
	for _, s := range add_news_update_db.GetAllServiceTypes() {
		if s == req.ServiceOpt {
			valid = true
			break
		}
	}
	if !valid {
		http.Error(w, "Invalid service option", http.StatusBadRequest)
		return
	}

	err := subscription_db.Subscribe(db.GetDB(), req.Email, req.FullName, req.ServiceOpt)
	if err != nil {
		http.Error(w, "Failed to subscribe", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Subscribed successfully",
	})
}

func handleUnsubscribe(w http.ResponseWriter, r *http.Request) {
	var req UnsubscribeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err := subscription_db.Unsubscribe(db.GetDB(), req.Email, req.ServiceOpt)
	if err != nil {
		http.Error(w, "Failed to unsubscribe", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Unsubscribed successfully",
	})
}

func handleGetServices(w http.ResponseWriter, r *http.Request) {
	services := add_news_update_db.GetAllServiceTypes()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"available_services": services,
	})
}
