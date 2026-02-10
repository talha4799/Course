package add_news_updates_routes

import (
	add_news_updates_controller "login/controllers/add_news_updates"
	"net/http"
)

func AddNewsUpdatesRoutes() {
	http.HandleFunc("/api/add_updates", add_news_updates_controller.AddNewsUpdates)
}
