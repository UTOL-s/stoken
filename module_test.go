package stoken

import (
	"errors"
	"testing"

	"github.com/ankorstore/yokai/config"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

// MockTokenClientFactory is a mock implementation of token.TokenClientFactory
type MockTokenClientFactory struct {
	tokenInitFunc func() error
}

// TokenInit implements token.TokenClientFactory
func (m *MockTokenClientFactory) TokenInit() error {
	return m.tokenInitFunc()
}

// createMockConfig creates a simple mock config for testing
func createMockConfig() *config.Config {
	v := viper.New()
	return &config.Config{Viper: v}
}

func TestNewTokenClientInit_Success(t *testing.T) {
	// Create a mock factory that returns nil error
	mockFactory := &MockTokenClientFactory{
		tokenInitFunc: func() error {
			return nil
		},
	}

	// Create the parameter struct with the mock factory
	// Note: Lifecycle is not used in NewTokenClientInit, so we can omit it
	param := FxTokenClientParam{
		Config:  createMockConfig(),
		Factory: mockFactory,
	}

	// Call the function under test
	err := NewTokenClientInit(param)

	// Assert that no error is returned
	assert.NoError(t, err)
}

func TestNewTokenClientInit_Error(t *testing.T) {
	// Create a mock factory that returns an error
	expectedErr := errors.New("token init error")
	mockFactory := &MockTokenClientFactory{
		tokenInitFunc: func() error {
			return expectedErr
		},
	}

	// Create the parameter struct with the mock factory
	// Note: Lifecycle is not used in NewTokenClientInit, so we can omit it
	param := FxTokenClientParam{
		Config:  createMockConfig(),
		Factory: mockFactory,
	}

	// Call the function under test
	err := NewTokenClientInit(param)

	// Assert that the expected error is returned
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
}

func TestModuleName(t *testing.T) {
	// Test that the ModuleName constant has the expected value
	assert.Equal(t, "stoken", ModuleName)
}

func TestFxSTokenModule(t *testing.T) {
	// Test that FxSTokenModule is not nil
	assert.NotNil(t, FxSTokenModule)
}
