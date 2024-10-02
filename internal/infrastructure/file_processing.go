package infrastructure

import (
	"encoding/json"
	"fmt"
	"os"
)

// LoadWordsFromFile loads categories and words from a JSON file.
func LoadWordsFromFile(filePath string) (map[string][]string, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	var categories map[string][]string
	err = json.NewDecoder(file).Decode(&categories)

	if err != nil {
		return nil, err
	}

	return categories, nil
}
