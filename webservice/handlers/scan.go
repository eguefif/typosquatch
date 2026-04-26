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

func Scan(w http.ResponseWriter, r *http.Request) {
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

	scanningJob := launchScan(domain)
	fmt.Fprintf(w, "scanning job: %d", scanningJob)
}

func launchScan(domain string) int {
	permutations := append(permutationengine.GetDomainPermutations(domain), []string{domain}...)
	go func() {
		results := checker.CheckTypoSquatting(permutations)
		fmt.Println("Result for job %d", 1)
		fmt.Println(results)
	}()
	return 1
}

func stripWWW(domain string) string {
	idx := strings.Index(domain, "www.")
	if idx != 0 {
		return domain
	}
	return domain[4:]
}
