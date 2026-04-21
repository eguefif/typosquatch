package permutationengine

import "iter"

func PermutationsGenerator(base_domain string) iter.Seq[string] {
	return func(yield func(string) bool) {
		for n := range len(base_domain) {
			domain := charDeletion(base_domain, n)
			if !yield(domain) {
				return
			}
		}
	}
}

func charDeletion(domain string, n int) string {
	return domain[:n] + domain[n+1:]
}
