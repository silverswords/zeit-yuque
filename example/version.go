package example

import (
	"fmt"
	"net/http"
)

// Version get the version of api.
func Version(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API Version 1.0-alpha")
}
