package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBanku(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Banku Suite")
}
