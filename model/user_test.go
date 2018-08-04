package model_test

import (
	"testing"

	"github.com/meetup/model"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	t.Run("user ok", func(t *testing.T) {
		_, err := model.NewUser("ricardo", 20)

		assert.NoError(t, err)
	})

	t.Run("user without name", func(t *testing.T) {
		_, err := model.NewUser("", 20)
		assert.Error(t, err)
	})

	t.Run("user with invalid age", func(t *testing.T) {
		_, err := model.NewUser("ricardo", 18)
		assert.Error(t, err)
	})
}
