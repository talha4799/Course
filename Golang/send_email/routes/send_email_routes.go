package send_email_routes

import (
	send_email_controller "login/controllers"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/send-email", send_email_controller.SendEmail)
}
