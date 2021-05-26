package crud_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/geometry-labs/go-service-template/crud"
	"github.com/geometry-labs/go-service-template/fixtures"
)

var (
	testFixtures  fixtures.Fixtures
	blockRawModel *crud.BlockRawModel
)

func NewBlockModel() *crud.BlockRawModel {
	dsn := crud.NewDsn("localhost", "5432", "postgres", "changeme", "test_db", "disable", "UTC")
	postgresConn, _ := crud.NewPostgresConn(dsn)
	testBlockRawModel := crud.NewBlockRawModel(postgresConn.GetConn())
	return testBlockRawModel
}

func Init(fixturesFile string) {
	//fixtures.ReadCurrentDir()
	testFixtures, _ = fixtures.LoadTestFixtures(fixturesFile)
	blockRawModel = NewBlockModel()
}

func TestCrud(t *testing.T) {
	RegisterFailHandler(Fail)
	Init(fixtures.Block_raws_fixture)
	RunSpecs(t, "Crud Suite")
}
