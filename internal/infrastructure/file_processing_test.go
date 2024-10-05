package infrastructure_test

import (
	"encoding/json"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	. "hangman-game/internal/infrastructure"
)

// Mock for HangmanReadCloser
type MockHangmanReadCloser struct {
	mock.Mock
}

func (m *MockHangmanReadCloser) Read(p []byte) (n int, err error) {
	args := m.Called(p)
	return args.Int(0), args.Error(1)
}

func (m *MockHangmanReadCloser) Close() error {
	args := m.Called()
	return args.Error(0)
}

// Test for OpenFile function
func TestOpenFile_Success(t *testing.T) {
	filePath := "test_file.json"
	// Create a temporary file for testing
	file, err := os.Create(filePath)
	assert.NoError(t, err)

	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Errorf("Error removing file: %v", err)
		}
	}(filePath) // Clean up after test

	// Close the file
	err = file.Close()
	assert.NoError(t, err)

	// Test OpenFile
	gotFile, err := OpenFile(filePath)
	assert.NoError(t, err)
	assert.NotNil(t, gotFile)
	defer func(gotFile HangmanReadCloser) {
		err := gotFile.Close()
		if err != nil {
			t.Errorf("Error closing file: %v", err)
		}
	}(gotFile) // Ensure the file is closed after test
}

func TestOpenFile_FileNotFound(t *testing.T) {
	filePath := "non_existent_file.json"
	gotFile, err := OpenFile(filePath)
	assert.Error(t, err)
	assert.Nil(t, gotFile)
}

// Test for DecodeCategoriesFromFile function
func TestDecodeCategoriesFromFile_Success(t *testing.T) {
	mockFile := new(MockHangmanReadCloser)

	// Create a sample JSON data
	categories := map[string][]string{
		"animals": {"cat", "dog"},
		"fruits":  {"apple", "banana"},
	}
	data, _ := json.Marshal(categories)

	// Mock the Read method to simulate reading JSON data
	mockFile.On("Read", mock.Anything).Return(len(data), nil).Run(func(args mock.Arguments) {
		copy(args.Get(0).([]byte), data) // Simulate writing data into buffer
	})
	mockFile.On("Close").Return(nil)

	gotCategories, err := DecodeCategoriesFromFile(mockFile)

	assert.NoError(t, err)
	assert.Equal(t, categories, gotCategories)
	mockFile.AssertExpectations(t)
}

func TestDecodeCategoriesFromFile_DecodeError(t *testing.T) {
	mockFile := new(MockHangmanReadCloser)
	mockFile.On("Close").Return(nil)

	// Provide invalid JSON
	mockFile.On("Read", mock.Anything).Return(0, errors.New("read error"))

	gotCategories, err := DecodeCategoriesFromFile(mockFile)

	assert.Error(t, err)
	assert.Nil(t, gotCategories)
	mockFile.AssertExpectations(t)
}

// Test for LoadWordsFromFile function
func TestLoadWordsFromFile_Success(t *testing.T) {
	filePath := "test_file.json"
	categories := map[string][]string{
		"animals": {"cat", "dog"},
		"fruits":  {"apple", "banana"},
	}
	data, _ := json.Marshal(categories)
	_ = os.WriteFile(filePath, data, 0644)
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Errorf("Error removing file: %v", err)
		}
	}(filePath) // Clean up after test

	gotCategories, err := LoadWordsFromFile(filePath)
	assert.NoError(t, err)
	assert.Equal(t, categories, gotCategories)
}

func TestLoadWordsFromFile_OpenFileError(t *testing.T) {
	filePath := "non_existent_file.json"
	gotCategories, err := LoadWordsFromFile(filePath)
	assert.Error(t, err)
	assert.Nil(t, gotCategories)
}
