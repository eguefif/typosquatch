package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"typosquatch/checker"
)

func (h *Handler) Result(w http.ResponseWriter, r *http.Request) {
	// TODO: Add logic to retrieve job and response result
	// Return pending if yet nothing
	splits := strings.Split(r.URL.Path[1:], "/")
	if len(splits) != 2 {
		http.NotFound(w, r)
		return
	}
	// TODO: Handle error: parsing and result
	jobId, _ := strconv.ParseInt(splits[1], 10, 64)
	result := h.PopResult(jobId)
	htmlResult := makeHtmlResult(result)
	fmt.Fprintf(w, "<p><h1>Result for job %d</h1></p><p> %s</p>", jobId, htmlResult)
}

func makeHtmlResult(results []checker.Result) string {
	html := ""
	for _, result := range results {
		if len(result.Records) > 0 || len(result.MxRecords) > 0 {
			html = html + "<p>"
			html = html + "<h2>" + result.Domain + "</h2>"
			html = html + "<h3>Records</h3>"
			html = html + strings.Join(result.Records, ", ")
			html = html + "<h3>MX Records</h3>"
			html = html + strings.Join(result.MxRecords, ", ")
			html = html + "</p>"
		}

	}
	return html
}
