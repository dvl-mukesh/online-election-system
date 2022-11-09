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

type UserDAO struct {
	Server     string
	Database   string
	Collection string
}

var Collection *mongo.Collection
var ctx = context.TODO()

func (e *UserDAO) Connect() {
	clientOptions := options.Client().ApplyURI(e.Server)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	Election_Collection = client.Database(e.Database).Collection(e.Collection)
}

func (e *UserDAO) Insert(User model.User) error {
	_, err := Election_Collection.InsertOne(ctx, User)

	if err != nil {
		fmt.Print(err.Error())
		return errors.New("unable to create new record")
	}

	return nil
}

func (e *UserDAO) FindById(id string) ([]*model.User, error) {
	var Users []*model.User

	cur, err := Election_Collection.Find(ctx, bson.D{primitive.E{Key: "_id", Value: id}})

	if err != nil {
		return Users, errors.New("unable to query db")
	}

	for cur.Next(ctx) {
		var e model.User

		err := cur.Decode(&e)

		if err != nil {
			return Users, err
		}

		Users = append(Users, &e)
	}

	if err := cur.Err(); err != nil {
		return Users, err
	}

	cur.Close(ctx)

	if len(Users) == 0 {
		return Users, mongo.ErrNoDocuments
	}

	return Users, nil
}

func (e *UserDAO) FindByEmailAndPassword(email, password string) ([]*model.User, error) {
	var Users []*model.User

	cur, err := Election_Collection.Find(ctx, bson.D{primitive.E{Key: "email", Value: email}, primitive.E{Key: "password", Value: password}})

	if err != nil {
		return Users, errors.New("unable to query db")
	}

	for cur.Next(ctx) {
		var e model.User

		err := cur.Decode(&e)

		if err != nil {
			return Users, err
		}

		Users = append(Users, &e)
	}

	if err := cur.Err(); err != nil {
		return Users, err
	}

	cur.Close(ctx)

	if len(Users) == 0 {
		return Users, mongo.ErrNoDocuments
	}

	return Users, nil
}

func (e *UserDAO) Delete(id string) error {
	filter := bson.D{primitive.E{Key: "_id", Value: id}}

	res, err := Election_Collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("no Users were deleted")
	}

	return nil
}

func (epd *UserDAO) Update(user model.User) error {
	filter := bson.D{primitive.E{Key: "_id", Value: user.ID}}

	update := bson.D{primitive.E{Key: "$set", Value: user}}

	e := &model.User{}
	return Election_Collection.FindOneAndUpdate(ctx, filter, update).Decode(e)
}
