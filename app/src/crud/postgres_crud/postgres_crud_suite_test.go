package postgres_crud_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPostgresCrud(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PostgresCrud Suite")
}
