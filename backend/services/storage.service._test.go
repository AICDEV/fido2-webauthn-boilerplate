package services_test

import (
	"git.eon-cds.de/repos/dlab/wad-fido2/backend/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

type StorageTestObject struct {
	ID   string
	Data int
}

var testData = StorageTestObject{
	ID:   "TEST_USER_ID",
	Data: 845625,
}

func TestSaveInUserStore(t *testing.T) {
	services.GetUserStorageInstance().Save(testData.ID, testData)
}

func TestGetFromUserStore(t *testing.T) {
	data, err := services.GetUserStorageInstance().Get(testData.ID)

	assert.Nil(t, err)
	assert.Equal(t, data.(StorageTestObject).ID, testData.ID)
}

func TestGetFromUserStoreWithNonExistingEntry(t *testing.T) {
	_, err := services.GetUserStorageInstance().Get("NULL")

	assert.NotNil(t, err)
}

func TestSaveInCredentialStore(t *testing.T) {
	services.GetCredentialStorageInstance().Save(testData.ID, testData)
}

func TestGetFromCredentialStore(t *testing.T) {
	data, err := services.GetCredentialStorageInstance().Get(testData.ID)

	assert.Nil(t, err)
	assert.Equal(t, data.(StorageTestObject).ID, testData.ID)
}

func TestGetFromCredentialStoreWithNonExistingEntry(t *testing.T) {
	_, err := services.GetCredentialStorageInstance().Get("NULL")

	assert.NotNil(t, err)
}
