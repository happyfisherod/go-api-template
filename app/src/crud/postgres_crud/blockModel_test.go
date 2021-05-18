package postgres_crud_test

import (
	"github.com/geometry-labs/app/crud/postgres_crud"
	"github.com/geometry-labs/app/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BlockModel", func() {

	var (
		block         *models.BlockRaw // predefined block
		dsn           string
		postgresConn  *postgres_crud.PostgresConn
		blockRawModel *postgres_crud.BlockRawModel
		found         *models.BlockRaw
	)

	var _ = BeforeSuite(func() {
		//block = &models.BlockRaw{
		//	Signature:        "8fIQfMRMEopVfckgTLPJosVsvVAU7ST6dZdPlFDJfiNO/nDt1xQswvip2maQZcNQrnR+awCf1nvzmek0VfqucwE=",
		//	ItemId:           "block_7204c15adec70ce6c3219defaae99137dfcf652b1748011be96bceea084f03b2",
		//	NextLeader:       "",
		//	TransactionCount: 1,
		//	Type:             "block",
		//	Version:          "0.1a",
		//	PeerId:           "hx547c6d5f7e80cd97b95cfefbaad919549a80831c",
		//	Number:           809,
		//	MerkleRootHash:   "be3883e4935576ce6a52355a26fd3ae15958e314db3243c1071b243cf312e9fb",
		//	ItemTimestamp:    "2018-03-28T07:34:35Z",
		//	Hash:             "7204c15adec70ce6c3219defaae99137dfcf652b1748011be96bceea084f03b2",
		//	ParentHash:       "f42aa479ef0c7913d6f5b151d5431254d4196571e86f3adca89d7c0439144f9a",
		//	Timestamp:        1522222475159807,
		//}
		block = &(*testCrudData.Blocks)[0]

		dsn = postgres_crud.NewDsn("localhost", "5432", "postgres", "changeme", "test_db", "disable", "UTC")
		postgresConn, _ = postgres_crud.NewPostgresConn(dsn)
		blockRawModel = postgres_crud.NewBlockRawModel(postgresConn.GetConn())

		_ = blockRawModel.Migrate()
		blockRawModel.Delete("Signature = ?", block.Signature)
	})

	Describe("blockModel with postgres", func() {
		Context("insert in block table", func() {
			It("predefined block insert", func() {
				blockRawModel.Create(block)
				found, _ := blockRawModel.FindOne("Signature = ?", block.Signature)
				Expect(found.Hash).To(Equal(block.Hash))
			})
		})
		Context("update in block table", func() {
			BeforeEach(func() {
				blockRawModel.Create(block)
			})
			It("predefined block update", func() {
				blockRawModel.Update(found, &models.BlockRaw{Type: "blockRaw"}, "Signature = ?", block.Signature)
				found, _ = blockRawModel.FindOne("Signature = ?", block.Signature)
				Expect(found.Type).To(Equal("blockRaw"))
			})
		})
		Context("delete in block table", func() {
			BeforeEach(func() {
				blockRawModel.Create(block)
			})
			It("predefined block delete", func() {
				blockRawModel.Delete("Signature = ?", block.Signature)
				found, _ = blockRawModel.FindOne("Signature = ?", block.Signature)
				Expect(found.Hash).To(Equal(""))
			})
		})
	})

})
