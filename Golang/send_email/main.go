package main

import (
	"fmt"
	send_email_routes "login/routes"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	fmt.Println("Hello This is Send Email API project")

	send_email_routes.RegisterRoutes()

	http.HandleFunc("/", home)

	port := ":8086"
	fmt.Printf("Server running at http://localhost%s\n", port)

	err = http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Send Email API")
}
