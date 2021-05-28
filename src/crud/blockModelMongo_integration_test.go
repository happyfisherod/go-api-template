//+build integration

package crud_test

import (
	"fmt"
	"github.com/geometry-labs/go-service-template/fixtures"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mongo Integration test", func() {
	testFixtures, _ := fixtures.LoadTestFixtures(fixtures.Block_raws_fixture)

	Describe("BlockModel with mongodb", func() {

		Context("Insert in block collection", func() {
			//testFixtures, _ = fixtures.LoadTestFixtures(fixtures.Block_raws_fixture) //To
			for _, fixture := range testFixtures {
				block := fixture.GetBlock(fixture.Input)
				It("insert in mongodb", func() {
					//nm := crud.NewMongoConn("mongodb://127.0.0.1:27017")

					x := blockRawModelMongo.GetMongoConn().ListAllDatabases()
					fmt.Println(x)

					//brmm := crud.NewBlockRawModelMongo(nm)
					//coll := brmm.SetCollectionHandle("icon_test", "contracts")
					//coll.InsertOne(nm.GetCtx(), bson.D{
					//	{Key: "title", Value: "The Polyglot Developer Podcast"},
					//	{Key: "author", Value: "Nic Raboy"},
					//})
					one, err := blockRawModelMongo.GetCollectionHandle().InsertOne(blockRawModelMongo.GetMongoConn().GetCtx(), block)

					if err != nil {
						Expect(1).To(Equal(0))
					}
					fmt.Println(one.InsertedID)
					Expect(1).To(Equal(1))
				}) // It
			} // For each fixture
		}) // Context "Insert in block collection"

	}) // Describe "BlockModel with mongodb"

}) // Describe "BlockModelMongo"
