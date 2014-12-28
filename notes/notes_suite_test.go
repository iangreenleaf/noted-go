package notes

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestNoted(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Noted Suite")
}
