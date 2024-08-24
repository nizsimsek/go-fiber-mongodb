package app

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/nizsimsek/go-fiber-mongodb/mocks/services"
	"github.com/nizsimsek/go-fiber-mongodb/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
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

	var FakeDataForHandler = []models.Todo{
		{Id: primitive.NewObjectID(), Title: "Title 1", Content: "Content 1"},
		{Id: primitive.NewObjectID(), Title: "Title 2", Content: "Content 2"},
		{Id: primitive.NewObjectID(), Title: "Title 3", Content: "Content 3"},
	}

	mockService.EXPECT().TodoGetAll().Return(FakeDataForHandler, nil)

	req := httptest.NewRequest("GET", "/api/todos", nil)

	res, _ := router.Test(req, 1)

	assert.Equal(t, 200, res.StatusCode)
}
