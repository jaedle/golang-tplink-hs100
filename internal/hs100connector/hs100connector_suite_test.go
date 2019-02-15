package hs100connector_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHs100Connector(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Hs100Connector Suite")
}
