package db

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindOne(collection string, query Query, decoded interface{}) error {
	err := bandmatchDB.Collection(collection).FindOne(context.TODO(), query).Decode(decoded)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}

		return err
	}

	return nil
}

func FindOneOrUpsert(collection string, query Query, update Query, decoded interface{}) error {
	opts := options.FindOneAndUpdate().SetUpsert(true)
	err := bandmatchDB.Collection(collection).FindOneAndUpdate(context.TODO(), query, update, opts).Decode(decoded)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}

		return err
	}

	return nil
}

func FindMany(collection string, query Query, decoded interface{}) error {
	cursor, err := bandmatchDB.Collection(collection).Find(context.TODO(), query)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}

		return err
	}

	defer cursor.Close(context.TODO())
	err = cursor.All(context.TODO(), decoded)
	if err != nil {
		return err
	}

	return nil
}

func InsertOne(collection string, document interface{}) (primitive.ObjectID, error) {
	res, err := bandmatchDB.Collection(collection).InsertOne(context.TODO(), document)
	if err != nil {
		return primitive.NilObjectID, err
	}
	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, err
	}

	return id, nil
}

func UpdateOne(collection string, filter Query, update Query) error {
	res, err := bandmatchDB.Collection(collection).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return fmt.Errorf("could not update document: %v", err)
	}

	if res.MatchedCount == 0 && res.ModifiedCount == 0 {
		return fmt.Errorf("did not update document for an unspecified reason")
	}

	return nil
}

func UpsertOne(collection string, filter Query, update Query) error {
	opts := options.Update().SetUpsert(true)
	res, err := bandmatchDB.Collection(collection).UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		return fmt.Errorf("could not update document: %v", err)
	}

	if res.MatchedCount == 0 && res.ModifiedCount == 0 {
		return fmt.Errorf("did not update document for an unspecified reason")
	}

	return nil
}

func DeleteOne(collection string, query Query) (int, error) {
	res, err := bandmatchDB.Collection(collection).DeleteOne(context.TODO(), query)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, nil
		}

		return -1, err
	}

	return int(res.DeletedCount), nil
}

func DeleteMany(collection string, query Query) (int, error) {
	res, err := bandmatchDB.Collection(collection).DeleteMany(context.TODO(), query)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, nil
		}

		return -1, err
	}

	return int(res.DeletedCount), nil
}

func getBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
