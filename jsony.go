package jsony

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

var ErrReadingFile = errors.New("error reading file")
var ErrWritingFile = errors.New("error writing file")
var ErrIncorrectJson = errors.New("incorrect JSON")
var ErrIncorrectStruct = errors.New("incorrect structure for JSON")

func Read(filename string, expectedStructure interface{}) error {
	// error common to all function
	var err error
	// red in the file
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrReadingFile, err)
	}
	// check for validity of the JSON data
	ok := json.Valid(fileContent)
	if !ok {
		return fmt.Errorf("%w", ErrIncorrectJson)
	}
	// unmarshal the data
	err = json.Unmarshal(fileContent, expectedStructure)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrIncorrectStruct, err)
	}

	return nil
}

func Write(data interface{}, filename string) error {
	var err error
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrIncorrectStruct, err)
	}
	err = os.WriteFile(filename, jsonData, 0600)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrWritingFile, err)
	}
	return nil
}