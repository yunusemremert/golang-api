package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-api/models"
	"time"
)

type TodoRepositoryDB struct {
	TodoCollection *mongo.Collection
}

type TodoRepository interface {
	Insert(todo models.Todo) (bool, error)
}

func (t TodoRepositoryDB) Insert(todo models.Todo) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.TodoCollection.InsertOne(ctx, todo)
	if result.InsertedID != nil || err != nil {
		err := errors.New("failed add")
		if err != nil {
			return false, err
		}

		return false, err
	}

	return true, nil
}

func NewTodoRepositoryDB(dbClient *mongo.Collection) TodoRepositoryDB {
	return TodoRepositoryDB{TodoCollection: dbClient}
}
