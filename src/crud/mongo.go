package crud

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"sync"
)

type MongoConn struct {
	client *mongo.Client
	ctx    context.Context
}

var mongoInstance *MongoConn
var mongoConnOnce sync.Once

func GetMongoConn() *MongoConn {
	mongoConnOnce.Do(func() {
		// TODO: create uri string from env variables
		uri := "mongodb://mongo:27017"
		client, err := mongo.NewClient(options.Client().ApplyURI(uri).SetAuth(options.Credential{
			AuthMechanism:           "",
			AuthMechanismProperties: nil,
			AuthSource:              "",
			Username:                "mongo",
			Password:                "changethis",
			PasswordSet:             true,
		}))
		if err != nil {
			log.Fatal("Cannot create a connection to mongodb", err)
		}
		//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		//defer cancel()
		ctx, _ := context.WithCancel(context.Background())
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal("Cannot connect to context for mongodb", err)
		}
		mongoInstance = &MongoConn{
			client: client,
			ctx:    ctx,
		}
	})
	return mongoInstance
}

func (m *MongoConn) GetClient() *mongo.Client {
	return m.client
}

func (m *MongoConn) GetCtx() context.Context {
	return m.ctx
}

func NewMongoConn(uri string) *MongoConn {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri).SetAuth(options.Credential{
		AuthMechanism:           "",
		AuthMechanismProperties: nil,
		AuthSource:              "",
		Username:                "mongo",
		Password:                "changethis",
		PasswordSet:             true,
	}))
	if err != nil {
		log.Fatal("Cannot create a connection to mongodb", err)
	}
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	ctx, _ := context.WithCancel(context.Background())
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Cannot connect to context for mongodb", err)
	}
	mongoInstance = &MongoConn{
		client: client,
		ctx:    ctx,
	}
	return mongoInstance
}

func (m *MongoConn) Close() error {
	err := m.client.Disconnect(m.ctx)
	if err != nil {
		log.Fatal("Cannot disconnect from mongodb", err)
	}
	return err
}

func (m *MongoConn) Ping() error {
	err := m.client.Ping(m.ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Cannot ping mongodb", err)
	}
	return err
}

func (m *MongoConn) ListAllDatabases() []string {
	databases, err := m.client.ListDatabaseNames(m.ctx, bson.M{})
	if err != nil {
		log.Fatal("Cannot List databases", err)
	}
	return databases
}

func (m *MongoConn) DatabaseHandle(database string) *mongo.Database {
	return m.client.Database(database)
}

// env variables
//
