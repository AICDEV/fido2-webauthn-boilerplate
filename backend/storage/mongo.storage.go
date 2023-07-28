package storage

import (
	"context"
	"log"
	"sync"

	"github.com/aicdev/fido2-webauthn-boilerplate/backend/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStorageInterface interface {
	Close()
	GetClient() *mongo.Client
}

type MongoStorage struct {
	client *mongo.Client
}

var (
	mOnce            sync.Once
	mStorageInstance MongoStorageInterface
)

func GetMongoStorageInstance() MongoStorageInterface {
	mOnce.Do(func() {
		envConfig := utils.ParseEnv()

		serverAPI := options.ServerAPI(options.ServerAPIVersion1)
		opts := options.Client().ApplyURI(envConfig.Mongo_Uri).SetServerAPIOptions(serverAPI)

		client, err := mongo.Connect(context.Background(), opts)

		if err != nil {
			log.Fatal(err.Error())
		}

		mStorageInstance = &MongoStorage{
			client: client,
		}
	})

	return mStorageInstance
}

func (ms *MongoStorage) GetClient() *mongo.Client {
	return ms.client
}

func (ms *MongoStorage) Close() {
	ms.client.Disconnect(context.Background())
}
