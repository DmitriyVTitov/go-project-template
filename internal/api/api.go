package api

import (
	"fmt"
	"html"
	"net/http"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	name := strings.Trim(r.URL.Path, "/")
	if name == "" {
		name = "Gopher"
	}

	fmt.Fprintf(w, "<!DOCTYPE html>\n")
	fmt.Fprintf(w, "Hello, %s!\n", html.EscapeString(name))
}
