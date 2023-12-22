package repository

import (
	"database/sql"
	"github.com/emrekursatt/JuniorProject/configs"
	"github.com/emrekursatt/JuniorProject/models"
	"log"
)

//go:generate mockgen -destination=../mocks/repository/mock_taskRepository.go -package=repository github.com/emrekursatt/JuniorProject/repository TaskRepository
type TaskRepositoryDB struct {
	db *sql.DB
}

type TaskRepository interface {
	Insert(task models.TaskEntity) (bool, error)
	Delete(code string) (bool, error)
	Update(code string, taskEntity models.TaskEntity) (bool, error)
	GetAllTasks() ([]models.TaskEntity, error)
	IsAlreadyTaskEntityExist(code string) (bool, error)
}

func NewTaskRepository(db *sql.DB) TaskRepositoryDB {
	return TaskRepositoryDB{db: db}
}

func (taskRepo *TaskRepositoryDB) Insert(task models.TaskEntity) (bool, error) {
	_, err := taskRepo.db.Exec("INSERT INTO  "+configs.TASK_TABLE+"(code ,title, description, status) VALUES(?, ?, ? , ?)", task.Code, task.Title, task.Description, task.Status)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (taskRepo *TaskRepositoryDB) Delete(code string) (bool, error) {
	_, err := taskRepo.db.Exec("DELETE FROM "+configs.TASK_TABLE+" WHERE code=?", code)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (taskRepo *TaskRepositoryDB) Update(code string, taskEntity models.TaskEntity) (bool, error) {

	query, err := taskRepo.db.Exec("UPDATE "+configs.TASK_TABLE+" SET code=? ,title=?, description=?, status=? WHERE code=?", taskEntity.Code, taskEntity.Title, taskEntity.Description, taskEntity.Status, code)
	println(query)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (taskRepo *TaskRepositoryDB) GetAllTasks() ([]models.TaskEntity, error) {
	var tasks []models.TaskEntity
	var task models.TaskEntity

	rows, err := taskRepo.db.Query("SELECT code , title,description , status FROM " + configs.TASK_TABLE + " ORDER BY id ASC")
	if err != nil {
		return tasks, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&task.Code, &task.Title, &task.Description, &task.Status)
		if err != nil {
			return tasks, err
		}

		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (taskRepo *TaskRepositoryDB) IsAlreadyTaskEntityExist(code string) (bool, error) {
	var count int
	query := "SELECT COUNT(*) FROM task_table WHERE code=?"
	err := taskRepo.db.QueryRow(query, code).Scan(&count)
	if err != nil {
		log.Println(err)
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil

}
