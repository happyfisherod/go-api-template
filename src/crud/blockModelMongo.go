package crud

import (
	"github.com/geometry-labs/go-service-template/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlockRawModelMongo struct {
	mongoConn *MongoConn
	model     *models.BlockRaw
	//databaseHandle *mongo.Database
	//collectionHandle *mongo.Collection
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
