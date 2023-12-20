package services

import (
	"errors"
	"github.com/emrekursatt/JuniorProject/dto"
	"github.com/emrekursatt/JuniorProject/models"
	"github.com/emrekursatt/JuniorProject/repository"
	"log"
	"strconv"
)

type DefaultUserService struct {
	Repo repository.UserRepository
}

type UserService interface {
	Insert(user models.UserEntity) (*dto.UserDTO, error)
	Delete(user models.UserEntity) (*dto.UserDTO, error)
	UpdatePassword(user models.UserEntity) (*dto.UserDTO, error)
	GetAllUsers() ([]models.UserEntity, error)
	Login(user models.UserEntity) (*dto.UserDTO, error)
}

func NewUserService(repo repository.UserRepository) DefaultUserService {
	return DefaultUserService{Repo: repo}
}

func (u DefaultUserService) Insert(user models.UserEntity) (*dto.UserDTO, error) {
	userDTO := dto.UserDTO{}
	var entityExist, err = u.Repo.IsAlreadyUserEntityExist(user)
	if entityExist {
		userDTO.Status = false
		log.Println(err)
		return &userDTO, errors.New("User Already Exist : " + user.UserName)

	} else if user.UserName == "" || user.Password == "" || user.Email == "" {
		userDTO.Status = false
		return &userDTO, errors.New("Username or Password or Email Not Valid : " + "User Name : " + user.UserName + " Password :" + user.Password + " Email :" + user.Email)
	} else if len(strconv.Itoa(user.PhoneNumber)) < 10 {
		userDTO.Status = false
		return &userDTO, errors.New("Please enter the telephone number with 10 digits without leading zeros.")
	}

	result, err := u.Repo.Insert(user)

	if err != nil || result == false {
		log.Println(err)
		userDTO.Status = false
		return &userDTO, err
	}

	userDTO.Status = true
	return &userDTO, nil
}

func (u DefaultUserService) Delete(user models.UserEntity) (*dto.UserDTO, error) {
	userDTO := dto.UserDTO{}

	var entityExist, err = u.Repo.IsAlreadyUserEntityExist(user)
	if entityExist == false {
		log.Println(err)
		userDTO.Status = false
		return &userDTO, errors.New("User Not Found")
	}

	result, err := u.Repo.Delete(user)

	if err != nil || result == false {
		log.Println(err)
		userDTO.Status = false
		return &userDTO, err
	}

	userDTO.Status = true
	return &userDTO, nil
}

func (u DefaultUserService) UpdatePassword(user models.UserEntity) (*dto.UserDTO, error) {
	userDTO := dto.UserDTO{}

	var entityExist, err = u.Repo.IsAlreadyUserEntityExist(user)
	if entityExist == false {
		userDTO.Status = false
		return &userDTO, errors.New("User Not Found")
	}

	result, err := u.Repo.UpdatePassword(user)

	if err != nil || result == false {
		userDTO.Status = false
		return &userDTO, err
	}

	userDTO.Status = true
	return &userDTO, nil
}

func (u DefaultUserService) GetAllUsers() ([]models.UserEntity, error) {
	var users []models.UserEntity

	users, err := u.Repo.GetAllUsers()

	if err != nil {
		log.Println(err)
		return users, err
	}

	return users, nil
}

func (u DefaultUserService) Login(user models.UserEntity) (*dto.UserDTO, error) {
	userDTO := dto.UserDTO{}

	var entityExist, err = u.Repo.Login(user)
	if entityExist == false {
		log.Println(err)
		userDTO.Status = false
		return &userDTO, errors.New("User Not Found")
	}
	userDTO.Status = true
	return &userDTO, nil
}
