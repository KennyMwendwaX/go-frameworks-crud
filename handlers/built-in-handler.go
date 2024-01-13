package handlers

import (
	"fmt"
	"net/http"
	"time"
)

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

// Add other handlers if needed
