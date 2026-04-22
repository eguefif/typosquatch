package main

import (
	"fmt"
	"os"
	"strings"
	"typosquatch/checker"
	"typosquatch/permutationengine"
	"typosquatch/validator"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: typosquatch --domain YOURDOMAIN.com")
		return
	}
	args := os.Args[1:]
	if !(len(args) == 2 && args[0] == "--domain" && validator.ValidateDomain(args[1])) {
		fmt.Println(args[1], "is not valid.")
	}
	domain := os.Args[2]
	domain = stripWWW(domain)
	permutations := permutationengine.GetDomainPermutations(domain)
	results := checker.CheckTypoSquatting(permutations)
	fmt.Println(results)
}

func stripWWW(domain string) string {
	idx := strings.Index(domain, "www.")
	if idx != 0 {
		return domain
	}
	return domain[4:]
}
