package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"crud_ql/graph/model"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	connectionString := goDotEnvVariable("MONGO_DB_URI")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return &DB{
		client: client,
	}
}

func (db *DB) GetUser(id string) *model.User {
	userCollec := db.client.Database("graphql-job-board").Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	var user model.User
	err := userCollec.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	return &user
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

func (db *DB) CreateUser(userInfo model.CreateUserInput) *model.User {
	userCollec := db.client.Database("graphql-job-board").Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	inserg, err := userCollec.InsertOne(ctx, bson.M{
		"username": userInfo.Username,
		"email":    userInfo.Email,
		"password": userInfo.Password,
	})

	if err != nil {
		log.Fatal(err)
	}

	insertedID := inserg.InsertedID.(primitive.ObjectID).Hex()
	returnUser := model.User{
		ID:       insertedID,
		Username: userInfo.Username,
		Email:    userInfo.Email,
	}
	return &returnUser
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
