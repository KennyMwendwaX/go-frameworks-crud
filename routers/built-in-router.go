package routers

import (
	"net/http"

	"github.com/kenny-mwendwa/go-restapi-crud/handlers"
)

func SetUpBuiltInRouter() {
	http.HandleFunc("/", handlers.Greet)
	// Add other routes if needed
}
