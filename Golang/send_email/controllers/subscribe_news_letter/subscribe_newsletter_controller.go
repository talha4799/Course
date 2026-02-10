package subscribe_news_letter_controller

import "net/http"

func SubscribeNewsletter(w http.ResponseWriter, r *http.Request) {
	var Subscription = "Newsletter Subscription"
	w.Write([]byte(Subscription))
}
