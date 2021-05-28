package crud_test

import (
	"github.com/geometry-labs/go-service-template/crud"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	blockRawModel *crud.BlockRawModel
	//testFixtures fixtures.Fixtures
)

func TestCrud(t *testing.T) {
	RegisterFailHandler(Fail)
	//
	blockRawModel = NewBlockModel()
	//testFixtures, _ = fixtures.LoadTestFixtures(fixtures.Block_raws_fixture)
	//
	RunSpecs(t, "Crud Suite")
}

func NewBlockModel() *crud.BlockRawModel {
	dsn := crud.NewDsn("localhost", "5432", "postgres", "changeme", "test_db", "disable", "UTC")
	postgresConn, _ := crud.NewPostgresConn(dsn)
	testBlockRawModel := crud.NewBlockRawModel(postgresConn.GetConn())
	return testBlockRawModel
}
