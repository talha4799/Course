package newsletter_controller

import (
	"encoding/json"
	"login/db"
	"login/db/newsletter_db"
	"login/db/subscription_db"
	services "login/utils"
	"net/http"
)

type CreateNewsletterRequest struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	ServiceType string `json:"service_type"`
	SendNow     bool   `json:"send_now"`
}

func Newsletter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createNewsletter(w, r)
	case http.MethodGet:
		getNewsletters(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func createNewsletter(w http.ResponseWriter, r *http.Request) {
	var req CreateNewsletterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	newsletter := &newsletter_db.Newsletter{
		Title:       req.Title,
		Content:     req.Content,
		ServiceType: req.ServiceType,
		Status:      "draft",
	}

	id, err := newsletter_db.CreateNewsletter(db.GetDB(), newsletter)
	if err != nil {
		http.Error(w, "Failed to create newsletter", http.StatusInternalServerError)
		return
	}

	// If send_now is true, send immediately using goroutines
	if req.SendNow {
		newsletter.ID = id
		go sendNewsletterAsync(newsletter)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":        "Newsletter created successfully",
		"id":             id,
		"send_initiated": req.SendNow,
	})
}

func sendNewsletterAsync(newsletter *newsletter_db.Newsletter) {
	database := db.GetDB()

	// Update status to sending
	newsletter_db.UpdateNewsletterStatus(database, newsletter.ID, "sending")

	// Get subscribers
	subscribers, err := subscription_db.GetSubscribersByService(database, newsletter.ServiceType)
	if err != nil {
		newsletter_db.UpdateNewsletterStatus(database, newsletter.ID, "failed")
		return
	}

	// Send emails
	emailService := services.NewEmailService()
	err = emailService.SendNewsletter(newsletter, subscribers)

	if err != nil {
		newsletter_db.UpdateNewsletterStatus(database, newsletter.ID, "failed")
	} else {
		newsletter_db.UpdateNewsletterStatus(database, newsletter.ID, "sent")
	}
}

func getNewsletters(w http.ResponseWriter, r *http.Request) {
	newsletters, err := newsletter_db.GetPendingNewsletters(db.GetDB())
	if err != nil {
		http.Error(w, "Failed to fetch newsletters", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newsletters)
}
