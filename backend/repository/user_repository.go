package repository

import (
	"context"
	"log"
	"time"

	"crud_ql/graph/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (db *DB) GetUserBySub(sub string) (*model.User, error) {
	userCollec := db.client.Database("graphql-job-board").Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{"sub": sub}
	var user model.User
	err := userCollec.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *DB) GetUser(id string) (*model.User, error) {
	userCollec := db.client.Database("graphql-job-board").Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	var user model.User
	err := userCollec.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *DB) GetUsers() []*model.User {
	userCollec := db.client.Database("graphql-job-board").Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var users []*model.User
	cursor, err := userCollec.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(context.TODO(), &users); err != nil {
		panic(err)
	}

	return users
}

func (db *DB) CreateUser(userInfo model.CreateUserInput, sub *string) (*model.User, error) {
	userCollec := db.client.Database("graphql-job-board").Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if sub != nil {
		userInfo.Sub = sub
	}
	inserg, err := userCollec.InsertOne(ctx, userInfo)

	if err != nil {
		return nil, err
	}

	insertedID := inserg.InsertedID.(primitive.ObjectID)
	return db.GetUser(insertedID.Hex())
}

func (db *DB) UpdateUser(userId string, userInfo model.UpdateUserInput) *model.User {
	userCollec := db.client.Database("graphql-job-board").Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	updateJobInfo := bson.M{}

	if userInfo.Username != nil {
		updateJobInfo["username"] = *userInfo.Username
	}

	if userInfo.Email != nil {
		updateJobInfo["email"] = *userInfo.Email
	}

	if userInfo.Password != nil {
		updateJobInfo["password"] = *userInfo.Password
	}

	_id, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": updateJobInfo}

	results := userCollec.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var user model.User

	if err := results.Decode(&user); err != nil {
		log.Fatal(err)
	}

	return &user
}

func (db *DB) DeleteUser(userId string) *model.DeleteUserResponse {
	userCollec := db.client.Database("graphql-job-board").Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.M{"_id": _id}
	_, err := userCollec.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	return &model.DeleteUserResponse{DeletedUserID: userId}
}
