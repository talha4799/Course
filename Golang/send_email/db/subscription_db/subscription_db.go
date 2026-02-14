package subscription_db

import (
	"database/sql"
	"fmt"
	"time"
)

type Subscription struct {
	ID         int64     `json:"id"`
	Email      string    `json:"email"`
	FullName   string    `json:"full_name"`
	ServiceOpt string    `json:"service_opt"` // Which service they subscribed to
	CreatedAt  time.Time `json:"created_at"`
	Active     bool      `json:"active"`
}

func InitTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS subscriptions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL,
		full_name TEXT,
		service_opt TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		active BOOLEAN DEFAULT 1,
		UNIQUE(email, service_opt)
	);`

	_, err := db.Exec(query)
	if err != nil {
		panic(fmt.Sprintf("Failed to create subscriptions table: %v", err))
	}
	fmt.Println("subscriptions table initialized")
}

func Subscribe(db *sql.DB, email, fullName, serviceOpt string) error {
	query := `
	INSERT INTO subscriptions (email, full_name, service_opt, active)
	VALUES (?, ?, ?, 1)
	ON CONFLICT(email, service_opt) DO UPDATE SET active = 1, full_name = excluded.full_name`

	_, err := db.Exec(query, email, fullName, serviceOpt)
	return err
}

func Unsubscribe(db *sql.DB, email, serviceOpt string) error {
	query := `UPDATE subscriptions SET active = 0 WHERE email = ? AND service_opt = ?`
	_, err := db.Exec(query, email, serviceOpt)
	return err
}

func GetSubscribersByService(db *sql.DB, serviceOpt string) ([]Subscription, error) {
	query := `SELECT id, email, full_name, service_opt, created_at, active 
	          FROM subscriptions WHERE service_opt = ? AND active = 1`

	rows, err := db.Query(query, serviceOpt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subs []Subscription
	for rows.Next() {
		var s Subscription
		err := rows.Scan(&s.ID, &s.Email, &s.FullName, &s.ServiceOpt, &s.CreatedAt, &s.Active)
		if err != nil {
			return nil, err
		}
		subs = append(subs, s)
	}
	return subs, nil
}

func GetAllActiveSubscribers(db *sql.DB) ([]Subscription, error) {
	query := `SELECT id, email, full_name, service_opt, created_at, active 
	          FROM subscriptions WHERE active = 1 ORDER BY service_opt`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subs []Subscription
	for rows.Next() {
		var s Subscription
		err := rows.Scan(&s.ID, &s.Email, &s.FullName, &s.ServiceOpt, &s.CreatedAt, &s.Active)
		if err != nil {
			return nil, err
		}
		subs = append(subs, s)
	}
	return subs, nil
}
