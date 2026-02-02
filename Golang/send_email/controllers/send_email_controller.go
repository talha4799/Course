package send_email_controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"os"
)

type EmailRequest struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

func SendEmail(w http.ResponseWriter, r *http.Request) {
	var req EmailRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	from := os.Getenv("SENDER_EMAIL")
	password := os.Getenv("SENDER_PASSWORD")

	smtpHost := os.Getenv("SMTP_SERVER")
	smtpPort := os.Getenv("SMTP_PORT")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	msg := []byte("Subject: " + req.Subject + "\n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" +
		req.Message)

	err = smtp.SendMail(
		smtpHost+":"+smtpPort,
		auth,
		from,
		[]string{req.To},
		msg,
	)

	if err != nil {
		http.Error(w, fmt.Sprintf("Email failed: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Email sent successfully"))
}
