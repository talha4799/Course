package subscribe_news_letter_routes

import (
	subscribe_news_letter_controller "login/controllers/subscribe_news_letter"
	"net/http"
)

func SubscribeNewsLetterRoutes() {
	http.HandleFunc("/api/subscribe", subscribe_news_letter_controller.SubscribeNewsletter)
}
