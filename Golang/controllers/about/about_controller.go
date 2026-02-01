package about_controller

import (
	"fmt"
	"net/http"
)

func About(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "<h1>Welcome to About Page!</h1>")
	fmt.Fprintf(w, "<p>Server is running successfully! This is About Page</p>")
	fmt.Fprintf(w, "<p>We founded this on January 29, 2026</p>")
}
