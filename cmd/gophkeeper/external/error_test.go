package external

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestError(t *testing.T) {
	err := NewError(errors.New("test"), ErrorTypeConnection)
	assert.Equal(t, "External error: Cannot connect to store", err.Error())

	err = NewError(errors.New("test"), ErrorTypeRegister)
	assert.Equal(t, "External error: Registration failed", err.Error())

	err = NewError(errors.New("test"), ErrorTypeLogin)
	assert.Equal(t, "External error: Login failed", err.Error())

	err = NewError(errors.New("test"), ErrorTypeSaveToStorage)
	assert.Equal(t, "External error: Save to storage failed", err.Error())

	err = NewError(errors.New("test"), ErrorTypeRestoreFromStorage)
	assert.Equal(t, "External error: Restore from storage failed", err.Error())

	err = NewError(errors.New("test"), "unknown label")
	assert.Equal(t, "External error: Unknown error", err.Error())
}
