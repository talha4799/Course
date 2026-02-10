package newsletter_routes

import (
	newsletter_controller "login/controllers/newsletter"
	"net/http"
)

func NewsletterRoutes() {
	http.HandleFunc("/api/newsletter", newsletter_controller.Newsletter)
}
