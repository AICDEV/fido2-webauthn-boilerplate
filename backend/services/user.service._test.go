package services_test

import (
	"testing"

	"github.com/aicdev/fido2-webauthn-boilerplate/backend/services"
	"github.com/stretchr/testify/assert"
)

var (
	testEmail       = "test@test.com"
	testDisplayName = "Test"
)

func TestNewUserInstance(t *testing.T) {
	user := services.GetUserServiceInstance().New(testEmail, testDisplayName)

	assert.Equal(t, user.Email, testEmail)
	assert.Equal(t, user.DisplayName, testDisplayName)
	assert.NotNil(t, user.ID)
	assert.Len(t, user.ID, 64)
}
