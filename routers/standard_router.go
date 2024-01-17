package routers

import (
	"net/http"

	"github.com/kenny-mwendwa/go-restapi-crud/handlers"
)

func StandardRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/users", handlers.GetUsers)

	return mux

}
