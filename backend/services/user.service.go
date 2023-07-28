package services

import (
	"crypto/rand"
	"sync"

	"github.com/aicdev/fido2-webauthn-boilerplate/backend/models"
)

type UserServiceInterface interface {
	New(email string, displayName string, firstName string, lastName string) *models.User
}

type UserService struct {
}

var (
	userServiceSyncOnce sync.Once
	userServiceInstance UserServiceInterface
)

func GetUserServiceInstance() UserServiceInterface {
	userServiceSyncOnce.Do(func() {
		userServiceInstance = &UserService{}
	})

	return userServiceInstance
}

func (us *UserService) New(email string, displayName string, firstName string, lastName string) *models.User {

	ub := make([]byte, 64)
	rand.Read(ub)

	user := &models.User{
		RegistrationID: ub,
		Email:          email,
		DisplayName:    displayName,
		Firstname:      firstName,
		Lastname:       lastName,
	}

	return user
}
