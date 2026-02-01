package home_controller

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// content(w)

	fmt.Fprintf(w, "<h1>Welcome to Go Server!</h1>")
	fmt.Fprintf(w, "<p>Server is running successfully!</p>")
	fmt.Fprintf(w, "<p>Try these endpoints:</p>")
	fmt.Fprintf(w, "<ul>")
	fmt.Fprintf(w, "<li><a href='/contact'>/Contact US</a></li>")
	fmt.Fprintf(w, "<li><a href='/about'>About</a></li>")
	fmt.Fprintf(w, "<li><a href='/static/'>/static/</a></li>")
	fmt.Fprintf(w, "</ul>")
}

// func content(w http.ResponseWriter) {
// 	fmt.Fprintf(w, `<h1>Welcome to Go Server!</h1>`)
// }
