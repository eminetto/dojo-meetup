package store_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/meetup/model"
	"github.com/meetup/store"
)

func TestStoreUser(t *testing.T) {
	s := store.Storage{}
	user, _ := model.NewUser("marcelo", 19)
	err := s.StoreUser(user)
	assert.Equal(t, s.GetLastUser(), user)
	assert.NoError(t, err)
}

func TestStoreUserCopy(t *testing.T) {
	s := store.Storage{}
	user, _ := model.NewUser("marcelo", 19)
	s.StoreUser(user)
	user.Name = "marcel"
	assert.NotEqual(t, s.GetLastUser(), user)
}

func TestStoreUserEmpty(t *testing.T) {
	var s store.Storage
	assert.Nil(t, s.GetLastUser())
}
