package services

import (
	"github.com/emrekursatt/JuniorProject/mocks/repository"
	"github.com/emrekursatt/JuniorProject/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var serviceUser UserService
var mockRepoUser *repository.MockUserRepository

var TestDataUser = []models.UserEntity{
	{UserName: "User 1", Password: "Passwordone", Email: "user1@gmail.com", PhoneNumber: 1234567890},
	{UserName: "User 2", Password: "Passwordone", Email: "user2@gmail.com", PhoneNumber: 1234567890},
}

func setupUser(t *testing.T) func() {
	ctx := gomock.NewController(t)
	defer ctx.Finish()

	mockRepoUser = repository.NewMockUserRepository(ctx)
	serviceUser = NewUserService(mockRepoUser)

	return func() {
		serviceUser = nil
		defer ctx.Finish()
	}

}

func TestUserService_Insert(t *testing.T) {
	teardown := setupUser(t)
	defer teardown()

	for i, _ := range TestDataUser {
		user := TestDataUser[i]
		mockRepoUser.EXPECT().IsAlreadyUserEntityExist(gomock.Any()).Return(false, nil).Times(1)
		mockRepoUser.EXPECT().Insert(user).Return(true, nil).Times(1)
		result, err := serviceUser.Insert(user)

		assert.NotEmpty(t, result)

		if err != nil {
			t.Error("Error")
		}
	}
}

func TestUserService_Delete(t *testing.T) {
	teardown := setupUser(t)
	defer teardown()

	for i, _ := range TestDataUser {
		user := TestDataUser[i]
		mockRepoUser.EXPECT().IsAlreadyUserEntityExist(gomock.Any()).Return(true, nil).Times(1)
		mockRepoUser.EXPECT().Delete(gomock.Any()).Return(true, nil).Times(1)
		result, err := serviceUser.Delete(user.UserName)

		assert.NotEmpty(t, result)

		if err != nil {
			t.Error("Error")
		}
	}
}

func TestUserService_Update(t *testing.T) {
	teardown := setupUser(t)
	defer teardown()

	for i, _ := range TestDataUser {
		user := TestDataUser[i]
		mockRepoUser.EXPECT().IsAlreadyUserEntityExist(gomock.Any()).Return(true, nil).Times(1)
		mockRepoUser.EXPECT().Update(gomock.Any(), gomock.Any()).Return(true, nil).Times(1)
		result, err := serviceUser.Update(user.UserName, user)

		assert.NotEmpty(t, result)

		if err != nil {
			t.Error("Error")
		}
	}
}

func TestUserService_GetAll(t *testing.T) {
	teardown := setupUser(t)
	defer teardown()

	mockRepoUser.EXPECT().GetAllUsers().Return(TestDataUser, nil).Times(1)
	result, err := serviceUser.GetAllUsers()

	assert.NotEmpty(t, result)

	if err != nil {
		t.Error("Error")
	}
}
