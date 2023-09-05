package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang-api/dto"
	"golang-api/models"
	"golang-api/repository"
)

//go:generate mockgen -destination=../mocks/service/mockTodoService.go -package=services golang-api/services TodoService
type DefaultTodoService struct {
	Repo repository.TodoRepository
}

type TodoService interface {
	TodoInsert(todo models.Todo) (*dto.TodoDTO, error)
	TodoGetAll() ([]models.Todo, error)
	TodoDelete(id primitive.ObjectID) (bool, error)
}

func (d DefaultTodoService) TodoInsert(todo models.Todo) (*dto.TodoDTO, error) {
	var res dto.TodoDTO

	if len(todo.Title) <= 2 {
		res.Status = false

		return &res, nil
	}

	result, err := d.Repo.Insert(todo)
	if err != nil || result == false {
		res.Status = false

		return &res, err
	}

	res = dto.TodoDTO{Status: result}

	return &res, nil
}

func (d DefaultTodoService) TodoGetAll() ([]models.Todo, error) {
	result, err := d.Repo.GetAll()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (d DefaultTodoService) TodoDelete(id primitive.ObjectID) (bool, error) {
	result, err := d.Repo.Delete(id)
	if err != nil || result == false {
		return false, err
	}

	return true, nil
}

func NewTodoService(Repo repository.TodoRepository) DefaultTodoService {
	return DefaultTodoService{Repo: Repo}
}
