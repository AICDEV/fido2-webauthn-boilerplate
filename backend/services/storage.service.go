package services

import (
	"errors"
	"sync"
)

const (
	TEMP_STORE = "TEMP"
	CRED_STORE = "CRED"
)

type StorageInterface interface {
	Save(id string, data interface{})
	Get(id string) (data interface{}, err error)
}

type Storage struct {
	mut   sync.Mutex
	store map[string]interface{}
}

var (
	userStorageSyncOnce       sync.Once
	credentialStorageSyncOnce sync.Once
	userStorageInstance       StorageInterface
	credentialStorageInstance StorageInterface
)

func GetUserStorageInstance() StorageInterface {
	userStorageSyncOnce.Do(func() {
		userStorageInstance = &Storage{
			store: map[string]interface{}{},
		}
	})

	return userStorageInstance
}

func GetCredentialStorageInstance() StorageInterface {
	credentialStorageSyncOnce.Do(func() {
		credentialStorageInstance = &Storage{
			store: map[string]interface{}{},
		}
	})

	return credentialStorageInstance
}

func (s *Storage) Save(id string, data interface{}) {
	s.mut.Lock()
	defer s.mut.Unlock()

	s.store[id] = data
}

func (s *Storage) Get(id string) (data interface{}, err error) {
	s.mut.Lock()
	defer s.mut.Unlock()

	val, exists := s.store[id]

	if exists {
		return val, nil
	}

	return nil, errors.New("entry does not exists")
}
