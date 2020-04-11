package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDkr(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dkr Suite")
}
