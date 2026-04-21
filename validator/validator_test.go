package validator

import (
	"fmt"
	"testing"
)

func TestValidateDomainValidDomains(t *testing.T) {
	domains := []string{"www.basicdomain.com", "basicdomain.com", "localhost"}
	for _, domain := range domains {
		if ValidateDomain(domain) == false {
			t.Errorf(`ValidateDomain(%s) == false, want match true`, domain)
		}
	}
}

func TestValidateDomainInvalidDomains(t *testing.T) {
	fmt.Println("Running invalid domains")
	domains := []string{".basicdomain.com", "basicdomaincom.", "localhost-", "-localhost", "localhost-.com"}
	for _, domain := range domains {
		if ValidateDomain(domain) == true {
			t.Errorf(`ValidateDomain(%s) == true, want match false`, domain)
		}
	}
}
