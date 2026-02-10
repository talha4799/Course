package routes_init

import (
	add_news_updates_routes "login/routes/add_news_updates"
	newsletter_routes "login/routes/newsletter"
	subscribe_news_letter_routes "login/routes/subscribe_news_letter"
)

func RegisterRoutes() {
	subscribe_news_letter_routes.SubscribeNewsLetterRoutes()
	newsletter_routes.NewsletterRoutes()
	add_news_updates_routes.AddNewsUpdatesRoutes()
}
