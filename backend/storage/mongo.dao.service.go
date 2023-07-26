package storage

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDaoInterface interface {
	Save(obj interface{}) (*mongo.InsertOneResult, error)
	GetById(id string, objRest interface{}) error
	Query(query bson.M, objRest interface{}) error
	QuerySingle(key string, value interface{}, objRest interface{}) error
	UpdateById(id string, objRest interface{}, objUpdate interface{}) error
	UpdateSingle(key string, value interface{}, objRest interface{}, objUpdate interface{}) error
	DeleteById(id string) error
}

type MongoDaoService struct {
	Database       string
	Collection     string
	StorageAdapter MongoStorageInterface
}

func (mds *MongoDaoService) Save(obj interface{}) (*mongo.InsertOneResult, error) {
	return mds.getCollectionReference().InsertOne(context.Background(), obj)
}

func (mds *MongoDaoService) GetById(id string, objRest interface{}) error {

	objectId, err := mds.stringToObjectId(id)

	if err != nil {
		return err
	}

	ctx, cancel := mds.getDAOContext()
	defer cancel()

	return mds.getCollectionReference().FindOne(
		ctx,
		bson.M{
			"_id": objectId,
		},
	).Decode(objRest)

}

func (mds *MongoDaoService) QuerySingle(key string, value interface{}, objRest interface{}) error {

	ctx, cancel := mds.getDAOContext()
	defer cancel()

	return mds.getCollectionReference().FindOne(
		ctx,
		bson.M{
			key: value,
		},
	).Decode(objRest)

}

func (mds *MongoDaoService) Query(query bson.M, objRest interface{}) error {

	ctx, cancel := mds.getDAOContext()
	defer cancel()

	cursor, err := mds.getCollectionReference().Find(
		ctx,
		query,
	)

	if err != nil {
		return err
	}

	defer cursor.Close(ctx)

	return cursor.All(
		ctx,
		objRest,
	)

}

func (mds *MongoDaoService) UpdateById(id string, objRest interface{}, objUpdate interface{}) error {

	ctx, cancel := mds.getDAOContext()
	defer cancel()

	objectId, err := mds.stringToObjectId(id)

	if err != nil {
		return err
	}

	return mds.getCollectionReference().FindOneAndUpdate(
		ctx,
		bson.M{
			"_id": objectId,
		},
		bson.M{
			"$set": objUpdate,
		},
	).Decode(objRest)

}

func (mds *MongoDaoService) UpdateSingle(key string, value interface{}, objRest interface{}, objUpdate interface{}) error {

	ctx, cancel := mds.getDAOContext()
	defer cancel()

	return mds.getCollectionReference().FindOneAndUpdate(
		ctx,
		bson.M{
			key: value,
		},
		bson.M{
			"$set": objUpdate,
		},
	).Decode(objRest)

}

func (mds *MongoDaoService) DeleteById(id string) error {

	ctx, cancel := mds.getDAOContext()
	defer cancel()

	objectId, err := mds.stringToObjectId(id)

	if err != nil {
		return err
	}

	_, err = mds.getCollectionReference().DeleteOne(
		ctx,
		bson.M{
			"_id": objectId,
		},
		nil,
	)

	return err
}

func (mds *MongoDaoService) stringToObjectId(id string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id)
}

func (mds *MongoDaoService) getCollectionReference() *mongo.Collection {
	return mds.StorageAdapter.GetClient().Database(mds.Database).Collection(mds.Collection)
}

func (mds *MongoDaoService) getCustomUpdateOptions() *options.FindOneAndUpdateOptions {
	return options.FindOneAndUpdate().SetReturnDocument(options.After)
}

func (mds *MongoDaoService) getDAOContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 15*time.Second)
}
