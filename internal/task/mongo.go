package task

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"square-service/pkg/logging"
)

type mongoDB struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

func (m *mongoDB) Create(ctx context.Context, dto *CreateTaskDTO) (string, error) {

	task := *dto
	res, err := m.collection.InsertOne(ctx, task)
	if err != nil {
		//m.logger.Error("failed to insertOne", zap.Error(err))
		return "", err
	}

	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		//m.logger.Error("failed to parse objId", zap.Error(err))
		//m.logger.Debug("")
		return "", err
	}

	//m.logger.Info("insertOne success", zap.String("id", id.Hex()))
	return id.Hex(), nil

}

func (m *mongoDB) FindAll(ctx context.Context) ([]*Task, error) {

	res, err := m.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var tasks []*Task
	if err = res.All(ctx, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (m *mongoDB) FindTask(ctx context.Context, id string) (*Task, error) {

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		//m.logger.Error("failed on serializing objectID",
		//	zap.String("", id),
		//	zap.Error(err),
		//)
		return nil, err
	}

	filter := bson.M{"_id": objectID}
	singleResult := m.collection.FindOne(ctx, filter)
	if err = singleResult.Err(); err != nil {
		//m.logger.Error("failed to get elem from collection", zap.Error(err))
		return nil, err
	}

	var task *Task
	if err = singleResult.Decode(&task); err != nil {
		//m.logger.Error("failed to decode", zap.Error(err))
		return nil, err
	}

	return task, nil

}

func (m *mongoDB) UpdateTask(ctx context.Context, task Task) (err error) {

	objectID, err := primitive.ObjectIDFromHex(task.ID)
	if err != nil {
		//m.logger.Error("failed to get objID from id", zap.String("id", task.ID), zap.Error(err))
		return
	}

	filter := bson.M{"_id": objectID}
	taskBytes, err := bson.Marshal(task)
	if err != nil {
		//m.logger.Error("failed to marshall task", zap.Error(err))
		return
	}

	var updateTaskObj bson.M
	err = bson.Unmarshal(taskBytes, &updateTaskObj)
	if err != nil {
		//m.logger.Error("failed to unmarshall task", zap.Error(err))
		return
	}

	delete(updateTaskObj, "_id")
	update := bson.M{"$set": updateTaskObj}

	updateResult, err := m.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		//m.logger.Error("failed to update task", zap.Error(err))
		return
	}

	if updateResult.MatchedCount == 0 {
		//m.logger.Info("fail to find task", zap.String("id", task.ID))
		return fmt.Errorf("fail to find task with id %s", task.ID)
	}

	if updateResult.ModifiedCount == 0 {
		//m.logger.Info("nothing to modify task", zap.String("id", task.ID))
	}

	//m.logger.Info("update succeeded",
	//	zap.Int("found", int(updateResult.MatchedCount)),
	//	zap.Int("updated", int(updateResult.ModifiedCount)))

	return

}

func (m *mongoDB) Delete(ctx context.Context, id string) (err error) {

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		//m.logger.Error("failed on serializing objectID",
		//	zap.String("", id),
		//	zap.Error(err),
		//)
		return
	}

	filter := bson.M{"_id": objectID}
	deleteResult, err := m.collection.DeleteOne(ctx, filter)
	if err != nil {
		//m.logger.Error("failed to get elem from collection", zap.Error(err))
		return
	}

	if deleteResult.DeletedCount == 0 {
		//m.logger.Error("failed to find elem in collection")
		return fmt.Errorf("find %d elems in collection", deleteResult.DeletedCount)
	}

	return

}

func NewMongoStorage(database *mongo.Database, collectionName string, logger *logging.Logger) Storage {
	//database.Collection("")
	return &mongoDB{
		collection: database.Collection(collectionName),
		logger:     logger,
	}
}
