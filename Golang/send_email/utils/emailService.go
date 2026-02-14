package services

import (
	"fmt"
	"log"
	"login/db/newsletter_db"
	subscription_db "login/db/subscription_db"
	"net/smtp"
	"os"
	"sync"
)

type EmailService struct {
	SMTPHost    string
	SMTPPort    string
	SenderEmail string
	SenderPass  string
}

func NewEmailService() *EmailService {
	return &EmailService{
		SMTPHost:    os.Getenv("SMTP_HOST"),
		SMTPPort:    os.Getenv("SMTP_PORT"),
		SenderEmail: os.Getenv("SENDER_EMAIL"),
		SenderPass:  os.Getenv("SENDER_PASSWORD"),
	}
}

// SendNewsletter sends emails to all subscribers using goroutines and channels
func (es *EmailService) SendNewsletter(newsletter *newsletter_db.Newsletter, subscribers []subscription_db.Subscription) error {
	if len(subscribers) == 0 {
		log.Println("No subscribers to send to")
		return nil
	}

	var wg sync.WaitGroup
	emailChan := make(chan subscription_db.Subscription, len(subscribers))
	resultChan := make(chan EmailResult, len(subscribers))

	// Worker pool - 5 concurrent workers
	workerCount := 5
	if len(subscribers) < workerCount {
		workerCount = len(subscribers)
	}

	// Start workers
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go es.emailWorker(newsletter, &wg, emailChan, resultChan)
	}

	// Send jobs to workers
	go func() {
		for _, sub := range subscribers {
			emailChan <- sub
		}
		close(emailChan)
	}()

	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Collect results
	var successCount, failCount int
	for result := range resultChan {
		if result.Success {
			successCount++
			log.Printf("✓ Email sent to %s", result.Email)
		} else {
			failCount++
			log.Printf("✗ Failed to send to %s: %v", result.Email, result.Error)
		}
	}

	log.Printf("Newsletter sending completed: %d success, %d failed", successCount, failCount)

	if failCount > 0 {
		return fmt.Errorf("partial failure: %d emails failed", failCount)
	}
	return nil
}

type EmailResult struct {
	Email   string
	Success bool
	Error   error
}

func (es *EmailService) emailWorker(newsletter *newsletter_db.Newsletter, wg *sync.WaitGroup,
	emailChan <-chan subscription_db.Subscription, resultChan chan<- EmailResult) {

	defer wg.Done()

	for sub := range emailChan {
		err := es.sendSingleEmail(newsletter, sub)
		resultChan <- EmailResult{
			Email:   sub.Email,
			Success: err == nil,
			Error:   err,
		}
	}
}

func (es *EmailService) sendSingleEmail(newsletter *newsletter_db.Newsletter, sub subscription_db.Subscription) error {
	// Compose email
	subject := fmt.Sprintf("Subject: %s\r\n", newsletter.Title)
	mime := "MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n"
	body := fmt.Sprintf("<h1>%s</h1><p>Hello %s,</p><div>%s</div>",
		newsletter.Title, sub.FullName, newsletter.Content)

	msg := []byte(subject + mime + body)

	// Send email
	addr := fmt.Sprintf("%s:%s", es.SMTPHost, es.SMTPPort)
	auth := smtp.PlainAuth("", es.SenderEmail, es.SenderPass, es.SMTPHost)

	to := []string{sub.Email}
	return smtp.SendMail(addr, auth, es.SenderEmail, to, msg)
}
