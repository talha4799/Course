package routes

import (
	"goproject/controllers/static"
	"net/http"
)

func RegisterRoutes() {
	// Serve HTML templates without .html extension
	http.HandleFunc("/", static.StaticController)

	// Serve assets only (CSS/JS/Images)
	assets := http.FileServer(http.Dir("./static/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", assets))

	// // http.HandleFunc("/", home_controller.Home)
	// http.HandleFunc("/contact", contact_us_controller.ContactUs)
	// http.HandleFunc("/about", about_controller.About)
}
