package utils_test

import (
	"testing"

	"git.eon-cds.de/repos/dlab/wad-fido2/backend/models"
	"git.eon-cds.de/repos/dlab/wad-fido2/backend/utils"
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
