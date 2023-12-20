package services

import (
	"github.com/emrekursatt/JuniorProject/mocks/repository"
	"github.com/emrekursatt/JuniorProject/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var service TaskService
var mockRepo *repository.MockTaskRepository

var TestData = []models.TaskEntity{
	{Code: "Code 1", Title: "Title 1", Description: "Desc1", Status: 1},
	{Code: "Code 2", Title: "Title 2", Description: "Desc2", Status: 0},
	{Code: "Code 3", Title: "Title 3", Description: "Desc3", Status: 0},
	{Code: "Code 4", Title: "Title 5", Description: "Desc4", Status: 1},
}

func setup(t *testing.T) func() {
	ctx := gomock.NewController(t)
	defer ctx.Finish()

	mockRepo = repository.NewMockTaskRepository(ctx)
	service = NewTaskService(mockRepo)

	return func() {
		service = nil
		defer ctx.Finish()
	}
}

func TestTaskService_GetAll(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	mockRepo.EXPECT().GetAllTasks().Return(TestData, nil).Times(1)
	result, err := service.GetAllTasks()

	assert.NotEmpty(t, result)

	if err != nil {
		t.Error("Error")
	}
}

func TestTaskService_Insert(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	for i, _ := range TestData {
		task := TestData[i]
		mockRepo.EXPECT().IsAlreadyTaskEntityExist(gomock.Any()).Return(false, nil).Times(1)
		mockRepo.EXPECT().Insert(task).Return(true, nil).Times(1)
		result, err := service.Insert(task)

		assert.NotEmpty(t, result)

		if err != nil {
			t.Error("Error")
		}
	}
}

func TestTaskService_Update(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	for i, _ := range TestData {
		task := TestData[i]
		mockRepo.EXPECT().IsAlreadyTaskEntityExist(gomock.Any()).Return(true, nil).Times(1)
		mockRepo.EXPECT().Update(task).Return(true, nil).Times(1)
		result, err := service.Update(task)

		assert.NotEmpty(t, result)

		if err != nil {
			t.Error("Error")
		}
	}
}

func TestTaskService_Delete(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	for i, _ := range TestData {
		task := TestData[i]
		mockRepo.EXPECT().IsAlreadyTaskEntityExist(gomock.Any()).Return(true, nil).Times(1)
		mockRepo.EXPECT().Delete(task).Return(true, nil).Times(1)
		result, err := service.Delete(task)

		assert.NotEmpty(t, result)

		if err != nil {
			t.Error("Error")
		}
	}
}
