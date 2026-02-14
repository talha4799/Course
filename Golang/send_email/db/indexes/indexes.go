package indexes

import (
	"database/sql"
	"fmt"
)

func RegisterIndexes(db *sql.DB) {
	indexes := []string{
		// News updates indexes
		`CREATE INDEX IF NOT EXISTS idx_news_service_type ON news_updates(service_type);`,
		`CREATE INDEX IF NOT EXISTS idx_news_created_at ON news_updates(created_at);`,
		`CREATE INDEX IF NOT EXISTS idx_news_date ON news_updates(year, month, day);`,

		// Newsletter indexes
		`CREATE INDEX IF NOT EXISTS idx_newsletter_service ON newsletters(service_type);`,
		`CREATE INDEX IF NOT EXISTS idx_newsletter_status ON newsletters(status);`,

		// Subscription indexes
		`CREATE INDEX IF NOT EXISTS idx_sub_service ON subscriptions(service_opt);`,
		`CREATE INDEX IF NOT EXISTS idx_sub_email ON subscriptions(email);`,
		`CREATE INDEX IF NOT EXISTS idx_sub_active ON subscriptions(active);`,
	}

	for _, idx := range indexes {
		_, err := db.Exec(idx)
		if err != nil {
			fmt.Printf("Warning: failed to create index: %v\n", err)
		}
	}
	fmt.Println("All indexes registered")
}
