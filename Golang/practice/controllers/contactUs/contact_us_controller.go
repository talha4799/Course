package contact_us_controller

import (
	"fmt"
	"net/http"
)

func ContactUs(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/contact" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "<h1>Welcome to Contact Us Page!</h1>")
	fmt.Fprintf(w, "<p>Server is running successfully! This is Contact Us Page</p>")
	fmt.Fprintf(w, "<p>We founded this on January 29, 2026</p>")
}
