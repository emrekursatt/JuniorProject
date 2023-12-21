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
	Update(code string) (bool, error)
	GetAllTasks() ([]models.TaskEntity, error)
	IsAlreadyTaskEntityExist(code string) (bool, error)
	FindTaskEntity(code string) (models.TaskEntity, error)
}

func NewTaskRepository(db *sql.DB) TaskRepositoryDB {
	return TaskRepositoryDB{db: db}
}

func (taskRepo *TaskRepositoryDB) Insert(task models.TaskEntity) (bool, error) {
	_, err := taskRepo.db.Exec("INSERT INTO  "+configs.TABLE_NAME+"(code ,title, description, status) VALUES(?, ?, ? , ?)", task.Code, task.Title, task.Description, task.Status)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (taskRepo *TaskRepositoryDB) Delete(code string) (bool, error) {
	_, err := taskRepo.db.Exec("DELETE FROM "+configs.TABLE_NAME+" WHERE code=?", code)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (taskRepo *TaskRepositoryDB) Update(code string) (bool, error) {
	var task = FindTaskEntity(code)

	_, err := taskRepo.db.Exec("UPDATE "+configs.TABLE_NAME+" SET code=? ,title=?, description=?, status=? WHERE code=?", task.Code, task.Title, task.Description, task.Status, task.Code)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (taskRepo *TaskRepositoryDB) FindTaskEntity(code string) (task models.TaskEntity, err error) {
	err = taskRepo.db.QueryRow("SELECT code , title,description , status FROM "+configs.TABLE_NAME+" WHERE code=?", code).Scan(&task.Code, &task.Title, &task.Description, &task.Status)
	if err != nil {
		log.Fatal(err)
		return task, err
	}
	return task, nil

}
{

}

func (taskRepo *TaskRepositoryDB) GetAllTasks() ([]models.TaskEntity, error) {
	var tasks []models.TaskEntity
	var task models.TaskEntity

	rows, err := taskRepo.db.Query("SELECT code , title,description , status FROM " + configs.TABLE_NAME + " ORDER BY id ASC")
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
		log.Fatal(err)
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil

}

