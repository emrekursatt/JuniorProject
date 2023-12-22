package repository

import (
	"database/sql"
	"github.com/emrekursatt/JuniorProject/configs"
	"github.com/emrekursatt/JuniorProject/models"
	"log"
)

type UserRepositoryDB struct {
	db *sql.DB
}

//go:generate mockgen -destination=../mocks/repository/mock_userRepository.go -package=repository github.com/emrekursatt/JuniorProject/repository UserRepository
type UserRepository interface {
	Insert(user models.UserEntity) (bool, error)
	Delete(userName string) (bool, error)
	Update(userName string, user models.UserEntity) (bool, error)
	GetAllUsers() ([]models.UserEntity, error)
	IsAlreadyUserEntityExist(userName string) (bool, error)
	Login(user models.UserEntity) (bool, error)
}

func NewUserRepository(db *sql.DB) UserRepositoryDB {
	return UserRepositoryDB{db: db}
}

func (userRepo *UserRepositoryDB) Insert(user models.UserEntity) (bool, error) {
	_, err := userRepo.db.Exec("INSERT INTO  "+configs.USER_TABLE+"(userName ,password, email , phone_number) VALUES(?, ?, ? , ?)", user.UserName, user.Password, user.Email, user.PhoneNumber)
	if err != nil {
		log.Println(err)
		return false, err
	}
	return true, nil
}

func (userRepo *UserRepositoryDB) Delete(userName string) (bool, error) {
	_, err := userRepo.db.Exec("DELETE FROM "+configs.USER_TABLE+" WHERE userName=?", userName)
	if err != nil {
		log.Println(err)
		return false, err
	}
	return true, nil
}

func (userRepo *UserRepositoryDB) Update(userName string, user models.UserEntity) (bool, error) {
	_, err := userRepo.db.Exec("UPDATE "+configs.USER_TABLE+" SET username=?  , password=? , email=? , phone_number=?  WHERE userName=?", user.UserName, user.Password, user.Email, user.PhoneNumber, userName)
	if err != nil {
		log.Println(err)
		return false, err
	}
	return true, nil
}

func (userRepo *UserRepositoryDB) GetAllUsers() ([]models.UserEntity, error) {
	var users []models.UserEntity
	var user models.UserEntity

	rows, err := userRepo.db.Query("SELECT username , password, email , phone_number FROM " + configs.USER_TABLE + " ORDER BY id ASC")
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.UserName, &user.Password, &user.Email, &user.PhoneNumber)
		if err != nil {
			log.Println(err)
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (userRepo *UserRepositoryDB) IsAlreadyUserEntityExist(userName string) (bool, error) {
	var count int
	query := "SELECT COUNT(*) FROM " + configs.USER_TABLE + " WHERE userName = ?"
	err := userRepo.db.QueryRow(query, userName).Scan(&count)
	if err != nil {
		log.Println(err)
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (userRepo *UserRepositoryDB) Login(user models.UserEntity) (bool, error) {
	var count int
	query := "SELECT COUNT(*) FROM " + configs.USER_TABLE + " WHERE userName = ? AND password = ?"
	err := userRepo.db.QueryRow(query, user.UserName, user.Password).Scan(&count)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}
