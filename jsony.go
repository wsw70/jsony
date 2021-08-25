package jsony

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

var ErrReadingFile = errors.New("error reading file")
var ErrIncorrectJson = errors.New("incorrect JSON")
var ErrIncorrectStruct = errors.New("incorrect structure for JSON")

func jsony(filename string, expectedStructure interface{}) error {
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
