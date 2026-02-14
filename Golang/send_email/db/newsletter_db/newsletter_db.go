package newsletter_db

import (
	"database/sql"
	"fmt"
	"time"
)

type Newsletter struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title"`
	Content     string     `json:"content"`
	ServiceType string     `json:"service_type"` // Which service this newsletter belongs to
	CreatedAt   time.Time  `json:"created_at"`
	SentAt      *time.Time `json:"sent_at,omitempty"`
	Status      string     `json:"status"` // draft, sending, sent, failed
}

func InitTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS newsletters (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		service_type TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		sent_at DATETIME,
		status TEXT DEFAULT 'draft'
	);`

	_, err := db.Exec(query)
	if err != nil {
		panic(fmt.Sprintf("Failed to create newsletters table: %v", err))
	}
	fmt.Println("newsletters table initialized")
}

func CreateNewsletter(db *sql.DB, newsletter *Newsletter) (int64, error) {
	query := `
	INSERT INTO newsletters (title, content, service_type, status)
	VALUES (?, ?, ?, ?)`

	result, err := db.Exec(query, newsletter.Title, newsletter.Content,
		newsletter.ServiceType, newsletter.Status)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func UpdateNewsletterStatus(db *sql.DB, id int64, status string) error {
	var query string
	var args []interface{}

	if status == "sent" {
		query = `UPDATE newsletters SET status = ?, sent_at = ? WHERE id = ?`
		args = []interface{}{status, time.Now(), id}
	} else {
		query = `UPDATE newsletters SET status = ? WHERE id = ?`
		args = []interface{}{status, id}
	}

	_, err := db.Exec(query, args...)
	return err
}

func GetPendingNewsletters(db *sql.DB) ([]Newsletter, error) {
	query := `SELECT id, title, content, service_type, created_at, sent_at, status 
	          FROM newsletters WHERE status = 'draft' ORDER BY created_at ASC`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var newsletters []Newsletter
	for rows.Next() {
		var n Newsletter
		var sentAt sql.NullTime
		err := rows.Scan(&n.ID, &n.Title, &n.Content, &n.ServiceType, &n.CreatedAt, &sentAt, &n.Status)
		if err != nil {
			return nil, err
		}
		if sentAt.Valid {
			n.SentAt = &sentAt.Time
		}
		newsletters = append(newsletters, n)
	}
	return newsletters, nil
}
