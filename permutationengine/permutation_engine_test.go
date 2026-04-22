package permutationengine

import "testing"

func TestPermutationsGenerator(t *testing.T) {
	domain := "example.com"
	expectedPermutations := []string{"xample.com", "eample.com", "exmple.com", "exaple.com", "examle.com", "exampe.com", "exampl.com", "examplecom", "example.om", "example.cm", "example.co"}
	i := 0
	for permutation := range PermutationsGenerator(domain) {
		if expectedPermutations[i] != permutation {
			t.Errorf(`%s != %s, want equal`, expectedPermutations[i], permutation)
		}
		i += 1
	}
}

func TestGetDomainPermutations(t *testing.T) {
	domain := "example.com"
	expectedPermutations := []string{"xample.com", "eample.com", "exmple.com", "exaple.com", "examle.com", "exampe.com", "exampl.com", "examplecom", "example.om", "example.cm", "example.co"}
	i := 0
	permutations := GetDomainPermutations(domain)
	for _, permutation := range permutations {
		if expectedPermutations[i] != permutation {
			t.Errorf(`%s != %s, want equal`, expectedPermutations[i], permutation)
		}
		i += 1
	}
}
