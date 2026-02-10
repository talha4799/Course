package newsletter_controller

import "net/http"

func Newsletter(w http.ResponseWriter, r *http.Request) {
	var NewsLetter = "Newsletter"
	w.Write([]byte(NewsLetter))
}
