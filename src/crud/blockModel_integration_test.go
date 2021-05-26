//+build integration

package crud_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/geometry-labs/go-service-template/models"
)

var _ = Describe("BlockModel", func() {

	Describe("blockModel with postgres", func() {

		Context("insert in block table", func() {
			for _, fixture := range testFixtures {
				block := fixture.GetBlock(fixture.Input)
				BeforeEach(func() {
					blockRawModel.Delete("Signature = ?", block.Signature)
				})
				It("predefined block insert", func() {
					blockRawModel.Create(block)
					found, _ := blockRawModel.FindOne("Signature = ?", block.Signature)
					Expect(found.Hash).To(Equal(block.Hash))
				})
			}
		})

		Context("update in block table", func() {
			for _, fixture := range testFixtures {
				block := fixture.GetBlock(fixture.Input)
				BeforeEach(func() {
					blockRawModel.Delete("Signature = ?", block.Signature)
					blockRawModel.Create(block)
				})
				It("predefined block update", func() {
					blockRawModel.Update(block, &models.BlockRaw{Type: "blockRaw"}, "Signature = ?", block.Signature)
					found, _ := blockRawModel.FindOne("Signature = ?", block.Signature)
					Expect(found.Type).To(Equal("blockRaw"))
				})
			}
		})

		Context("delete in block table", func() {
			for _, fixture := range testFixtures {
				block := fixture.GetBlock(fixture.Input)
				BeforeEach(func() {
					blockRawModel.Delete("Signature = ?", block.Signature)
					blockRawModel.Create(block)
				})
				It("predefined block delete", func() {
					blockRawModel.Delete("Signature = ?", block.Signature)
					found, _ := blockRawModel.FindOne("Signature = ?", block.Signature)
					Expect(found.Hash).To(Equal(""))
				})
			}
		})
	})

})
