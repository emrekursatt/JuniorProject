package services

import (
	"errors"
	"github.com/emrekursatt/JuniorProject/dto"
	"github.com/emrekursatt/JuniorProject/models"
	"github.com/emrekursatt/JuniorProject/repository"
	"log"
)

//go:generate mockgen -destination=../mocks/service/mock_taskService.go -package=services github.com/emrekursatt/JuniorProject/services TaskService
type DefaultTaskService struct {
	Repo repository.TaskRepository
}

type TaskService interface {
	Insert(task models.TaskEntity) (*dto.TaskDTO, error)
	Delete(code string) (*dto.TaskDTO, error)
	Update(code string, task models.TaskEntity) (*dto.TaskDTO, error)
	GetAllTasks() ([]models.TaskEntity, error)
}

func NewTaskService(repo repository.TaskRepository) DefaultTaskService {
	return DefaultTaskService{Repo: repo}
}

func (t DefaultTaskService) Insert(task models.TaskEntity) (*dto.TaskDTO, error) {
	taskDTO := dto.TaskDTO{}

	var entityExist, err = t.Repo.IsAlreadyTaskEntityExist(task.Code)
	if entityExist {
		log.Println("Task Code Already Exist : " + task.Code)
		taskDTO.Status = false
		return &taskDTO, errors.New("Task Code Already Exist")

	} else if task.Code == "" || task.Title == "" {
		log.Println("Code or Title Not Valid : " + "Title : " + task.Title + " Code :" + task.Code)
		taskDTO.Status = false
		return &taskDTO, errors.New("Code or Title Not Valid : " + "Title : " + task.Title + " Code :" + task.Code)
	}

	result, err := t.Repo.Insert(task)

	if err != nil || result == false {
		taskDTO.Status = false
		return &taskDTO, err
	}

	log.Println("Task Insert Succesful : " + task.Code)
	taskDTO.Status = true
	return &taskDTO, nil

}

func (t DefaultTaskService) Delete(code string) (*dto.TaskDTO, error) {
	taskDTO := dto.TaskDTO{}

	var entityExist, err = t.Repo.IsAlreadyTaskEntityExist(code)
	if entityExist == false {
		log.Println("Task Code Not Found : " + code)
		taskDTO.Status = false
		return &taskDTO, errors.New("Task Code Not Found")

	}

	result, err := t.Repo.Delete(code)
	if err != nil || result == false {
		taskDTO.Status = false
		return &taskDTO, err
	}

	log.Println("Task Delete Succesful : " + code)
	taskDTO.Status = true
	return &taskDTO, nil

}

func (t DefaultTaskService) Update(code string, task models.TaskEntity) (*dto.TaskDTO, error) {
	taskDTO := dto.TaskDTO{}

	var entityExist, err = t.Repo.IsAlreadyTaskEntityExist(code)
	if entityExist == false {
		log.Println("Task Code Not Found : " + code)
		taskDTO.Status = false
		return &taskDTO, errors.New("Task Code Not Found")

	}

	result, err := t.Repo.Update(code, task)
	if err != nil || result == false {
		taskDTO.Status = false
		return &taskDTO, err
	}

	log.Println("Task Update Succesful : " + code)
	taskDTO.Status = true
	return &taskDTO, nil

}

func (t DefaultTaskService) GetAllTasks() ([]models.TaskEntity, error) {
	tasks, err := t.Repo.GetAllTasks()
	taskDto := dto.TaskDTO{}
	if err != nil {
		taskDto.Status = false
		return tasks, err
	}

	return tasks, nil
}
