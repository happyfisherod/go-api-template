package crud

import (
	"github.com/geometry-labs/app/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

var block = &models.BlockRaw{
	Signature:        "8fIQfMRMEopVfckgTLPJosVsvVAU7ST6dZdPlFDJfiNO/nDt1xQswvip2maQZcNQrnR+awCf1nvzmek0VfqucwE=",
	ItemId:           "block_7204c15adec70ce6c3219defaae99137dfcf652b1748011be96bceea084f03b2",
	NextLeader:       "",
	TransactionCount: 1,
	Type:             "block",
	Version:          "0.1a",
	PeerId:           "hx547c6d5f7e80cd97b95cfefbaad919549a80831c",
	Number:           809,
	MerkleRootHash:   "be3883e4935576ce6a52355a26fd3ae15958e314db3243c1071b243cf312e9fb",
	ItemTimestamp:    "2018-03-28T07:34:35Z",
	Hash:             "7204c15adec70ce6c3219defaae99137dfcf652b1748011be96bceea084f03b2",
	ParentHash:       "f42aa479ef0c7913d6f5b151d5431254d4196571e86f3adca89d7c0439144f9a",
	Timestamp:        1522222475159807,
}

func TestBlockModel(t *testing.T) {
	//_ = exec.Command("createdb", "-p", "5432", "-h", "localhost", "-U", "postgres", "-e", "test_db")
	dsn := "host=localhost user=postgres password=changeme dbname=test_db port=5432 sslmode=disable TimeZone=UTC"
	newDsn := NewDsn("localhost", "5432", "postgres", "changeme", "test_db", "disable", "UTC")
	assert.Equal(t, newDsn, dsn)

	postgresConn, _ := NewPostgresConn(newDsn)
	blockRawModel := NewBlockRawModel(postgresConn.conn)

	_ = blockRawModel.Migrate()
	blockRawModel.Delete("Signature = ?", block.Signature)

	// test insert
	blockRawModel.Create(block)
	found, _ := blockRawModel.FindOne("Signature = ?", "8fIQfMRMEopVfckgTLPJosVsvVAU7ST6dZdPlFDJfiNO/nDt1xQswvip2maQZcNQrnR+awCf1nvzmek0VfqucwE=")
	assert.Equal(t, block.Signature, found.Signature)

	// test update
	blockRawModel.Update(found, &models.BlockRaw{Type: "blockRaw"}, "Signature = ?", "8fIQfMRMEopVfckgTLPJosVsvVAU7ST6dZdPlFDJfiNO/nDt1xQswvip2maQZcNQrnR+awCf1nvzmek0VfqucwE=")
	found, _ = blockRawModel.FindOne("Signature = ?", "8fIQfMRMEopVfckgTLPJosVsvVAU7ST6dZdPlFDJfiNO/nDt1xQswvip2maQZcNQrnR+awCf1nvzmek0VfqucwE=")
	assert.Equal(t, "blockRaw", found.Type)

	// test delete
	blockRawModel.Delete("Signature = ?", block.Signature)
	found, _ = blockRawModel.FindOne("Signature = ?", "")

}
