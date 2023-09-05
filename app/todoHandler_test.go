package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
	services "golang-api/mocks/service"
	"golang-api/models"
	"net/http/httptest"
	"testing"
)

var td TodoHandler
var mockService *services.MockTodoService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)

	mockService = services.NewMockTodoService(ctrl)

	td = TodoHandler{mockService}

	return func() {
		defer ctrl.Finish()
	}
}

func TestTodoHandler_GetAllTodo(t *testing.T) {
	trd := setup(t)
	defer trd()

	router := fiber.New()

	router.Get("/api/todos", td.GetAllTodo)

	var FakeData = []models.Todo{
		{primitive.NewObjectID(), "Title 1", "Content 1"},
		{primitive.NewObjectID(), "Title 2", "Content 2"},
		{primitive.NewObjectID(), "Title 3", "Content 3"},
	}

	mockService.EXPECT().TodoGetAll().Return(FakeData, nil)

	req := httptest.NewRequest("GET", "/api/todos", nil)

	resp, _ := router.Test(req, 1)

	assert.Equal(t, 200, resp.StatusCode)
}
