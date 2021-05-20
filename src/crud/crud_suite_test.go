package crud_test

import (
	"github.com/geometry-labs/app/crud"
	"github.com/geometry-labs/app/fixtures"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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
	testFixtures = fixtures.LoadTestFixtures(fixturesFile)
	blockRawModel = NewBlockModel()
}

func TestCrud(t *testing.T) {
	RegisterFailHandler(Fail)
	Init(fixtures.Block_raws_fixture)
	RunSpecs(t, "Crud Suite")
}
