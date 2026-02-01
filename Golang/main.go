package main

import (
	"fmt"
	"goproject/routes"
	"net/http"
)

func main() {

	routes.RegisterRoutes()

	port := ":8086"

	fmt.Println("Routes registered successfully")
	fmt.Printf("Server running at http://localhost%s\n", port)
	fmt.Println("Available pages:")
	fmt.Println("  http://localhost:8086/")
	fmt.Println("  http://localhost:8086/about")
	fmt.Println("  http://localhost:8086/contact")

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
