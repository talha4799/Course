package main

import (
	"fmt"
	"log"
	"login/db"
	"login/routes_init"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file, using defaults")
	}

	fmt.Println("🚀 Newsletter System Starting...")

	// Initialize database (creates tables, indexes, WAL mode)
	db.Initialize()

	fmt.Println("✅ Database initialized with WAL mode (.db file)")

	// Register routes
	log.Println("Registering routes...")
	routes_init.RegisterRoutes()

	port := ":8086"
	fmt.Printf("📧 Newsletter API running at http://localhost%s\n", port)
	fmt.Println("Available endpoints:")
	fmt.Println("  POST   /api/news         - Add news update")
	fmt.Println("  POST   /api/newsletter   - Create newsletter")
	fmt.Println("  GET    /api/newsletter   - List pending newsletters")
	fmt.Println("  POST   /api/subscribe    - Subscribe to newsletter")
	fmt.Println("  DELETE /api/subscribe    - Unsubscribe")
	fmt.Println("  GET    /api/subscribe    - List available services")

	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Server error: %v\n", err)
	}
}
