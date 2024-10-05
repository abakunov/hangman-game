package infrastructure

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func OpenFile(filePath string) (HangmanReadCloser, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func DecodeCategoriesFromFile(file HangmanReadCloser) (map[string][]string, error) {
	defer func(file io.ReadCloser) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	var categories map[string][]string
	err := json.NewDecoder(file).Decode(&categories)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func LoadWordsFromFile(filePath string) (map[string][]string, error) {
	file, err := OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	return DecodeCategoriesFromFile(file)
}
