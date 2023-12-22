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
	Delete(userName string) (*dto.UserDTO, error)
	Update(userName string, user models.UserEntity) (*dto.UserDTO, error)
	GetAllUsers() ([]models.UserEntity, error)
	Login(user models.UserEntity) (*dto.UserDTO, error)
}

func NewUserService(repo repository.UserRepository) DefaultUserService {
	return DefaultUserService{Repo: repo}
}

func (u DefaultUserService) Insert(user models.UserEntity) (*dto.UserDTO, error) {
	userDTO := dto.UserDTO{}
	number := strconv.Itoa(user.PhoneNumber)
	var entityExist, err = u.Repo.IsAlreadyUserEntityExist(user.UserName)
	if entityExist {
		userDTO.Status = false
		log.Println(err)
		return &userDTO, errors.New("User Already Exist : " + user.UserName)

	} else if user.UserName == "" || user.Password == "" || user.Email == "" {
		userDTO.Status = false
		log.Println("Username or Password or Email Not Valid : " + "User Name : " + user.UserName + " Password :" + user.Password + " Email :" + user.Email)
		return &userDTO, errors.New("Username or Password or Email Not Valid : " + "User Name : " + user.UserName + " Password :" + user.Password + " Email :" + user.Email)
	} else if len(strconv.Itoa(user.PhoneNumber)) != 10 || number[0] == '0' {
		userDTO.Status = false
		log.Println("Please enter the telephone number with 10 digits without leading zeros.")
		return &userDTO, errors.New("Please enter the telephone number with 10 digits without leading zeros.")
	}

	result, err := u.Repo.Insert(user)

	if err != nil || result == false {
		log.Println(err)
		userDTO.Status = false
		return &userDTO, err
	}
	log.Println("User Create Succesful : " + user.UserName)
	userDTO.Status = true
	return &userDTO, nil
}

func (u DefaultUserService) Delete(userName string) (*dto.UserDTO, error) {
	userDTO := dto.UserDTO{}

	var entityExist, err = u.Repo.IsAlreadyUserEntityExist(userName)
	if entityExist == false {
		log.Println(errors.New("User Not Found"))
		userDTO.Status = false
		return &userDTO, errors.New("User Not Found")
	}

	result, err := u.Repo.Delete(userName)

	if err != nil || result == false {
		log.Println(err)
		userDTO.Status = false
		return &userDTO, err
	}

	userDTO.Status = true
	return &userDTO, nil
}

func (u DefaultUserService) Update(userName string, user models.UserEntity) (*dto.UserDTO, error) {
	userDTO := dto.UserDTO{}

	var entityExist, err = u.Repo.IsAlreadyUserEntityExist(userName)
	if entityExist == false {
		userDTO.Status = false
		log.Println("Please enter the telephone number with 10 digits without leading zeros.")
		return &userDTO, errors.New("User Not Found")
	} else if len(strconv.Itoa(user.PhoneNumber)) != 10 {
		userDTO.Status = false
		log.Println("Please enter the telephone number with 10 digits without leading zeros.")
		return &userDTO, errors.New("Please enter the telephone number with 10 digits without leading zeros.")
	}

	result, err := u.Repo.Update(userName, user)

	if err != nil || result == false {
		userDTO.Status = false
		return &userDTO, err
	}

	userDTO.Status = true
	return &userDTO, nil
}

//go:generate mockgen -destination=../mocks/service/mock_userService.go -package=services github.com/emrekursatt/JuniorProject/services UserService

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
