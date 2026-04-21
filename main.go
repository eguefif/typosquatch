package main

import (
	"fmt"
	"os"
	"typosquatch/permutationengine"
	"typosquatch/validator"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: typosquatch --domain YOURDOMAIN.com")
		return
	}
	args := os.Args[1:]
	if len(args) == 2 && args[0] == "--domain" && validator.ValidateDomain(args[1]) {
		fmt.Println(args[1], "is a valid domain.")
	} else {
		fmt.Println(args[1], "is not valid.")
	}
	domain := os.Args[2]
	for t := range permutationengine.PermutationsGenerator(domain) {
		fmt.Println(t)
	}
}
