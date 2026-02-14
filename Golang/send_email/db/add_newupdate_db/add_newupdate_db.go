package add_news_update_db

import (
	"database/sql"
	"fmt"
	"time"
)

// Service options constants
const (
	TechUpdates          = "Tech Updates"
	SportsUpdates        = "Sports Updates"
	EntertainmentUpdates = "Entertainment Updates"
	HealthUpdates        = "Health Updates"
	BusinessUpdates      = "Business Updates"
	ScienceUpdates       = "Science Updates"
	WorldUpdates         = "World Updates"
	LocalUpdates         = "Local Updates"
	SpecialOccasions     = "Special Occasions"
	StrangeEvents        = "Strange events and accidents"
)

type NewsUpdate struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	ServiceType string    `json:"service_type"` // One of the 10 service options
	Day         int       `json:"day"`
	Month       int       `json:"month"`
	Year        int       `json:"year"`
	CreatedAt   time.Time `json:"created_at"`
}

func InitTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS news_updates (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		service_type TEXT NOT NULL,
		day INTEGER,
		month INTEGER,
		year INTEGER,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := db.Exec(query)
	if err != nil {
		panic(fmt.Sprintf("Failed to create news_updates table: %v", err))
	}
	fmt.Println("news_updates table initialized")
}

func AddNewsUpdate(db *sql.DB, update *NewsUpdate) (int64, error) {
	query := `
	INSERT INTO news_updates (title, content, service_type, day, month, year, created_at)
	VALUES (?, ?, ?, ?, ?, ?, ?)`

	result, err := db.Exec(query, update.Title, update.Content, update.ServiceType,
		update.Day, update.Month, update.Year, time.Now())
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func GetNewsUpdatesByService(db *sql.DB, serviceType string) ([]NewsUpdate, error) {
	query := `SELECT id, title, content, service_type, day, month, year, created_at 
	FROM news_updates WHERE service_type = ? ORDER BY created_at DESC`

	rows, err := db.Query(query, serviceType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var updates []NewsUpdate
	for rows.Next() {
		var u NewsUpdate
		err := rows.Scan(&u.ID, &u.Title, &u.Content, &u.ServiceType, &u.Day, &u.Month, &u.Year, &u.CreatedAt)
		if err != nil {
			return nil, err
		}
		updates = append(updates, u)
	}
	return updates, nil
}

func GetAllServiceTypes() []string {
	return []string{
		TechUpdates, SportsUpdates, EntertainmentUpdates, HealthUpdates,
		BusinessUpdates, ScienceUpdates, WorldUpdates, LocalUpdates,
		SpecialOccasions, StrangeEvents,
	}
}
