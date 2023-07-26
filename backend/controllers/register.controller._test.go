package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"git.eon-cds.de/repos/dlab/wad-fido2/backend/controllers"
	"git.eon-cds.de/repos/dlab/wad-fido2/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegisterControllerSignUpPayloader(t *testing.T) {

	prepareTestEnv()

	testPayload := models.ActionSignUpRequest{
		Email:       "wrong",
		Firstname:   "",
		Lastname:    "",
		DisplayName: "",
	}

	testPayloadBytes, _ := json.Marshal(testPayload)

	testRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(testRecorder)

	req, _ := http.NewRequest(http.MethodPost, "", bytes.NewBuffer(testPayloadBytes))
	req.Header.Set("content-type", "application/json")

	ctx.Request = req

	registerControllerInstance := controllers.GetRegisterControllerInstance()
	registerControllerInstance.Begin(ctx)

	assert.Equal(t, testRecorder.Result().StatusCode, http.StatusBadRequest)
}

func prepareTestEnv() {
	os.Setenv("WAD_PORT", "16666")
	os.Setenv("WAD_KEYPATH", "no exists")
	os.Setenv("WAD_MONGO_URI", "mongodb://localhost")
	os.Setenv("WAD_REDIS_HOST", "no exists")
	os.Setenv("WAD_REDIS_PORT", "no exists")
	os.Setenv("WAD_REDIS_PASSWORD", "no exists")
}
