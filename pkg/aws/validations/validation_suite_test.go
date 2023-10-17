package validations

import (
    "testing"

    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
)

func TestValidations(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Validations Suite")
}
