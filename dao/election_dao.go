package dao

import (
	"context"
	"errors"
	"fmt"
	"log"

	"online-election-system/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ElectionDAO struct {
	Server     string
	Database   string
	Collection string
}

var Election_Collection *mongo.Collection
var election_ctx = context.TODO()

func (e *ElectionDAO) Connect() {
	clientOptions := options.Client().ApplyURI(e.Server)
	client, err := mongo.Connect(election_ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(election_ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	Election_Collection = client.Database(e.Database).Collection(e.Collection)
}

func (e *ElectionDAO) Insert(Election model.Election) error {
	_, err := Election_Collection.InsertOne(election_ctx, Election)

	if err != nil {
		fmt.Print(err.Error())
		return errors.New("unable to create new record")
	}

	return nil
}

func (e *ElectionDAO) FindById(id string) ([]*model.Election, error) {
	var Elections []*model.Election

	cur, err := Election_Collection.Find(election_ctx, bson.D{primitive.E{Key: "_id", Value: id}})

	if err != nil {
		return Elections, errors.New("unable to query db")
	}

	for cur.Next(election_ctx) {
		var e model.Election

		err := cur.Decode(&e)

		if err != nil {
			return Elections, err
		}

		Elections = append(Elections, &e)
	}

	if err := cur.Err(); err != nil {
		return Elections, err
	}

	cur.Close(election_ctx)

	if len(Elections) == 0 {
		return Elections, mongo.ErrNoDocuments
	}

	return Elections, nil
}

func (e *ElectionDAO) FindByEmailAndPassword(email, password string) ([]*model.Election, error) {
	var Elections []*model.Election

	cur, err := Election_Collection.Find(election_ctx, bson.D{primitive.E{Key: "email", Value: email}, primitive.E{Key: "password", Value: password}})

	if err != nil {
		return Elections, errors.New("unable to query db")
	}

	for cur.Next(election_ctx) {
		var e model.Election

		err := cur.Decode(&e)

		if err != nil {
			return Elections, err
		}

		Elections = append(Elections, &e)
	}

	if err := cur.Err(); err != nil {
		return Elections, err
	}

	cur.Close(election_ctx)

	if len(Elections) == 0 {
		return Elections, mongo.ErrNoDocuments
	}

	return Elections, nil
}

func (e *ElectionDAO) Delete(id string) error {
	filter := bson.D{primitive.E{Key: "_id", Value: id}}

	res, err := Election_Collection.DeleteOne(election_ctx, filter)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("no Elections were deleted")
	}

	return nil
}

func (epd *ElectionDAO) Update(Election model.Election) error {
	filter := bson.D{primitive.E{Key: "_id", Value: Election.ID}}

	update := bson.D{primitive.E{Key: "$set", Value: Election}}

	e := &model.Election{}
	return Election_Collection.FindOneAndUpdate(election_ctx, filter, update).Decode(e)
}
