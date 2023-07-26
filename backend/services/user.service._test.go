package services_test

import (
	"git.eon-cds.de/repos/dlab/wad-fido2/backend/services"
	"github.com/stretchr/testify/assert"
	"testing"
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
