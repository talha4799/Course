package add_news_updates_controller

import "net/http"

func AddNewsUpdates(w http.ResponseWriter, r *http.Request) {
	var newsUpdates = "News Updates"
	w.Write([]byte(newsUpdates))
}
