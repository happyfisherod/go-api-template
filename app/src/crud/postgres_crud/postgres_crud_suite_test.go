package postgres_crud_test

import (
	test "github.com/geometry-labs/app/test_related"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var testCrudData test.TestCrudData

func TestPostgresCrud(t *testing.T) {
	RegisterFailHandler(Fail)
	testCrudData.Blocks = test.LoadTestBlocks("../../../test_fixtures/crud/postgres_crud/block_raws.json")
	RunSpecs(t, "PostgresCrud Suite")
}

//var _ = BeforeSuite(func() {
//	Expect(testCrudData.Blocks).NotTo(BeEmpty())
//})
