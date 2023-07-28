package utils_test

import (
	"testing"

	"github.com/aicdev/fido2-webauthn-boilerplate/backend/models"
	"github.com/aicdev/fido2-webauthn-boilerplate/backend/utils"
	"github.com/stretchr/testify/assert"
)

func TestRedisStoreRealizer(t *testing.T) {
	test := &models.User{
		ID:          "SOME-TEST-ID",
		Email:       "test@test.com",
		DisplayName: "TestUser",
	}

	enc, err := utils.EncodeForRedisStore(test)

	assert.Nil(t, err)

	testRest := &models.User{}

	err = utils.DecodeFromRedisStore(string(enc), testRest)

	assert.Nil(t, err)
	assert.Equal(t, test.Email, testRest.Email)
}
