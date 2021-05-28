package crud_test

import (
	"github.com/geometry-labs/go-service-template/crud"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	blockRawModel      *crud.BlockRawModel
	blockRawModelMongo *crud.BlockRawModelMongo
	//testFixtures fixtures.Fixtures
)

func TestCrud(t *testing.T) {
	RegisterFailHandler(Fail)
	//
	blockRawModel = NewBlockModel()
	blockRawModelMongo = NewBlockModelMongo()
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

func NewBlockModelMongo() *crud.BlockRawModelMongo {
	mongoConn := crud.NewMongoConn("mongodb://127.0.0.1:27017")
	blockRawModelMongo := crud.NewBlockRawModelMongo(mongoConn)
	_ = blockRawModelMongo.SetCollectionHandle("icon_test", "contracts")
	return blockRawModelMongo
}
