package utils_test

import (
	"net/http"
	"testing"

	"github.com/aicdev/fido2-webauthn-boilerplate/backend/utils"
	"github.com/stretchr/testify/assert"
)

func TestEmailFilter(t *testing.T) {
	validInput := "test@test.com"

	filtered := utils.FilterStorageQuery(validInput)

	assert.Equal(t, validInput, filtered)
}

func TestTokenExtractionExpected(t *testing.T) {
	testReq, _ := http.NewRequest(http.MethodGet, "localhost", nil)
	testReq.Header.Set("Authorization", "Bearer FOOBARTOKEN")

	token, err := utils.FilterTokenStringFromAuthorizationHeader(testReq)

	assert.Nil(t, err)
	assert.Equal(t, token, "FOOBARTOKEN")
}

func TestTokenExtractionBadBearer(t *testing.T) {
	testReq, _ := http.NewRequest(http.MethodGet, "localhost", nil)
	testReq.Header.Set("Authorization", "FOOBARTOKEN")

	token, err := utils.FilterTokenStringFromAuthorizationHeader(testReq)

	assert.NotNil(t, err)
	assert.Equal(t, token, "")
}

func TestTokenExtractionBadToken(t *testing.T) {
	testReq, _ := http.NewRequest(http.MethodGet, "localhost", nil)
	testReq.Header.Set("Authorization", "Bearer ")

	token, err := utils.FilterTokenStringFromAuthorizationHeader(testReq)
	assert.NotNil(t, err)
	assert.Equal(t, token, "")

}
