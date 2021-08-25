package jsony

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

var errReadingFile = errors.New("error reading file")
var errWritingFile = errors.New("error writing file")
var errIncorrectJson = errors.New("incorrect JSON")
var errIncorrectStruct = errors.New("incorrect structure for JSON")

// Read JSON data from a file (filename) and push it into the container (expectedStructure)
func Read(filename string, expectedStructure interface{}) error {
	// error common to all function
	var err error
	// red in the file
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("%w: %v", errReadingFile, err)
	}
	// check for validity of the JSON data
	ok := json.Valid(fileContent)
	if !ok {
		return fmt.Errorf("%w", errIncorrectJson)
	}
	// unmarshal the data
	err = json.Unmarshal(fileContent, expectedStructure)
	if err != nil {
		return fmt.Errorf("%w: %v", errIncorrectStruct, err)
	}

	return nil
}

// Write JSON data into a file (filename)
func Write(data interface{}, filename string) error {
	var err error
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("%w: %v", errIncorrectStruct, err)
	}
	err = os.WriteFile(filename, jsonData, 0600)
	if err != nil {
		return fmt.Errorf("%w: %v", errWritingFile, err)
	}
	return nil
}
