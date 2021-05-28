package crud

import (
	"github.com/geometry-labs/go-service-template/models"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

type BlockRawModelMongo struct {
	mongoConn *MongoConn
	model     *models.BlockRaw
	//databaseHandle *mongo.Database
	//collectionHandle *mongo.Collection
}

var blockRawModelMongoInstance *BlockRawModelMongo
var blockRawModelMongoOnce sync.Once

func GetBlockRawModelMongo() *BlockRawModelMongo {
	blockRawModelMongoOnce.Do(func() {
		blockRawModelMongoInstance = &BlockRawModelMongo{
			mongoConn: GetMongoConn(),
			model:     &models.BlockRaw{},
		}
	})
	return blockRawModelMongoInstance
}

func NewBlockRawModelMongo(conn *MongoConn) *BlockRawModelMongo {
	blockRawModelMongoInstance := &BlockRawModelMongo{
		mongoConn: conn,
		model:     &models.BlockRaw{},
	}
	return blockRawModelMongoInstance
}

func (b *BlockRawModelMongo) CollectionHandle(database string, collection string) *mongo.Collection {
	return b.mongoConn.DatabaseHandle(database).Collection(collection)
}
