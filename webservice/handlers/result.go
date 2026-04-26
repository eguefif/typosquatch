package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

func Result(w http.ResponseWriter, r *http.Request) {
	splits := strings.Split(r.URL.Path[1:], "/")
	if len(splits) != 2 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Result for job: %s", splits[1])
}
