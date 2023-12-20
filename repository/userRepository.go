package repository

import (
	"database/sql"
	"github.com/emrekursatt/JuniorProject/configs"
	"github.com/emrekursatt/JuniorProject/models"
)

type UserRepositoryDB struct {
	db *sql.DB
}

type UserRepository interface {
	Insert(user models.UserEntity) (bool, error)
	Delete(user models.UserEntity) (bool, error)
	UpdatePassword(user models.UserEntity) (bool, error)
	GetAllUsers() ([]models.UserEntity, error)
	IsAlreadyUserEntityExist(user models.UserEntity) (bool, error)
}

func NewUserRepository(db *sql.DB) UserRepositoryDB {
	return UserRepositoryDB{db: db}
}

func (userRepo *UserRepositoryDB) Insert(user models.UserEntity) (bool, error) {
	_, err := userRepo.db.Exec("INSERT INTO  "+configs.USER+"(userName ,password, email , phone_number) VALUES(?, ?, ? , ?)", user.UserName, user.Password, user.Email, user.PhoneNumber)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (userRepo *UserRepositoryDB) Delete(user models.UserEntity) (bool, error) {
	_, err := userRepo.db.Exec("DELETE FROM "+configs.USER+" WHERE userName=?", user.UserName)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (userRepo *UserRepositoryDB) UpdatePassword(user models.UserEntity) (bool, error) {
	_, err := userRepo.db.Exec("UPDATE "+configs.USER+" SET password=? WHERE userName=?", user.Password, user.UserName)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (userRepo *UserRepositoryDB) GetAllUsers() ([]models.UserEntity, error) {
	var users []models.UserEntity
	var user models.UserEntity

	rows, err := userRepo.db.Query("SELECT userName , password, email FROM " + configs.USER + " ORDER BY id ASC")
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.UserName, &user.Password, &user.Email)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (userRepo *UserRepositoryDB) IsAlreadyUserEntityExist(user models.UserEntity) (bool, error) {
	var count int
	query := "SELECT COUNT(*) FROM " + configs.USER + " WHERE userName = ?"
	err := userRepo.db.QueryRow(query, user.UserName).Scan(&count)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}
