package example

import (
	"fmt"
	"net/http"
	"time"
)

// Greeting is a naive handler for testing.
func Greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go on Now 2.0, current time is: %v", time.Now())
}
