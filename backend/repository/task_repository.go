package repository

import (
	"context"
	"log"
	"time"

	"crud_ql/graph/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *DB) GetTask(id string) *model.Task {
	taskCollec := db.client.Database("graphql-job-board").Collection("task")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	var task model.Task
	err := taskCollec.FindOne(ctx, filter).Decode(&task)
	if err != nil {
		log.Fatal(err)
	}

	// Fetch owner/user associated with this task
	owner := db.GetUser(*task.OwnerID)
	task.Owner = owner

	return &task
}

func (db *DB) GetTasks() []*model.Task {
	taskCollec := db.client.Database("graphql-job-board").Collection("task")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var tasks []*model.Task
	cursor, err := taskCollec.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(context.TODO(), &tasks); err != nil {
		panic(err)
	}

	// Fetch owner/user associated with this task
	for _, task := range tasks {
		owner := db.GetUser(*task.OwnerID)
		task.Owner = owner
	}

	return tasks
}

func (db *DB) CreateTask(taskInfo model.CreateTaskInput) *model.Task {
	taskCollec := db.client.Database("graphql-job-board").Collection("task")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	inserg, err := taskCollec.InsertOne(ctx, taskInfo)

	if err != nil {
		log.Fatal(err)
	}

	_id := inserg.InsertedID.(primitive.ObjectID)

	user := db.GetUser(taskInfo.OwnerID)

	return &model.Task{
		ID:                     _id.Hex(),
		Owner:                  user,
		Description:            taskInfo.Description,
		Category:               taskInfo.Category,
		TaskRequirements:       taskInfo.TaskRequirements,
		Location:               taskInfo.Location,
		Budget:                 taskInfo.Budget,
		SpecificSkillsRequired: taskInfo.SpecificSkillsRequired,
		Urgency:                taskInfo.Urgency,
		Priority:               taskInfo.Priority,
		Status:                 taskInfo.Status,
	}
}

func (db *DB) UpdateTask(id string, taskInfo model.UpdateTaskInput) *model.Task {
	taskCollec := db.client.Database("graphql-job-board").Collection("task")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": taskInfo}

	_, err := taskCollec.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}

	return db.GetTask(id)
}

func (db *DB) DeleteTask(id string) *model.DeleteTaskResponse {
	taskCollec := db.client.Database("graphql-job-board").Collection("task")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}

	_, err := taskCollec.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	return &model.DeleteTaskResponse{
		DeletedTaskID: id,
	}
}
