package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestForce(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Force Suite")
}
