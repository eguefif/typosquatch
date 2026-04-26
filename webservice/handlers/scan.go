package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"typosquatch/checker"
	"typosquatch/permutationengine"
	"typosquatch/validator"
)

// TODO: Create a context for the APP
// Get the next job id and returns it

func (h *Handler) Scan(w http.ResponseWriter, r *http.Request) {
	splits := strings.Split(r.URL.Path[1:], "/")
	if len(splits) != 2 {
		http.NotFound(w, r)
		return
	}

	// TODO: Should decode base64 in url

	domain := splits[1]
	domain = stripWWW(domain)
	if !validator.ValidateDomain(domain) {
		fmt.Fprintf(w, "Invalid domain: %s", domain)
		return
	}

	scanningJobId := h.launchScan(domain)
	fmt.Fprintf(w, "scanning job: %d", scanningJobId)
}

func (h *Handler) launchScan(domain string) int64 {
	permutations := append(permutationengine.GetDomainPermutations(domain), []string{domain}...)
	jobId := h.AddJob()
	go func() {
		results := checker.CheckTypoSquatting(permutations)
		h.AddResult(jobId, results)
		fmt.Println("Result for job %d", jobId)
		fmt.Println(results)
	}()
	return jobId
}

func stripWWW(domain string) string {
	idx := strings.Index(domain, "www.")
	if idx != 0 {
		return domain
	}
	return domain[4:]
}
