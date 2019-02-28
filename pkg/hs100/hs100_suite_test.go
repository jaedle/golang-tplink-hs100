package hs100_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHs100(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Hs100 Suite")
}
