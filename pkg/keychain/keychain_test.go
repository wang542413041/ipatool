package keychain

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestKeychain(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Keychain Suite")
}
