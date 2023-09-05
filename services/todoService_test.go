package services

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
	"golang-api/mocks/repository"
	"golang-api/models"
	"testing"
)

var mockRepo *repository.MockTodoRepository
var service TodoService

var FakeData = []models.Todo{
	{primitive.NewObjectID(), "Title 1", "Content 1"},
	{primitive.NewObjectID(), "Title 2", "Content 2"},
	{primitive.NewObjectID(), "Title 3", "Content 3"},
}

func setup(t *testing.T) func() {
	ct := gomock.NewController(t)
	defer ct.Finish()

	mockRepo = repository.NewMockTodoRepository(ct)
	service = NewTodoService(mockRepo)

	return func() {
		service = nil

		defer ct.Finish()
	}
}

func TestDefaultTodoService_TodoGetAll(t *testing.T) {
	td := setup(t)
	defer td()

	mockRepo.EXPECT().GetAll().Return(FakeData, nil)

	result, err := service.TodoGetAll()
	if err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, result)
}

func TestDefaultTodoService_TodoInsert(t *testing.T) {
	td := setup(t)
	defer td()

	todo := models.Todo{
		Id:      primitive.ObjectID{},
		Title:   "Title 4",
		Content: "Content 4",
	}

	mockRepo.EXPECT().Insert(todo).Return(true, nil)

	result, err := service.TodoInsert(todo)
	if err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, result)
}

func TestDefaultTodoService_TodoDelete(t *testing.T) {
	td := setup(t)
	defer td()

	id := FakeData[0].Id

	mockRepo.EXPECT().Delete(id).Return(true, nil)

	result, err := service.TodoDelete(id)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, result, true)
}
