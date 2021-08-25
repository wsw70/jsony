package jsony

import (
	"errors"
	"testing"
)

var correctStruct = struct {
	Hello string
}{}

var incorrectStruct = struct {
	Hello int
}{}

func TestNoFile(t *testing.T) {
	err := Read("nonexistingfile.json", &correctStruct)
	if !errors.Is(err, errReadingFile) {
		t.Errorf("inexisting file not detected: %v", err)
	} else {
		t.Logf("👍: %v", err)
	}
}

func TestIncorrectJson(t *testing.T) {
	err := Read("testJsonIncorrect.json", &correctStruct)
	if !errors.Is(err, errIncorrectJson) {
		t.Errorf("incorrect JSON not detected: %v", err)
	} else {
		t.Logf("👍: %v", err)
	}
}

func TestCorrectJsonIncorrectStruct(t *testing.T) {
	err := Read("testJsonCorrect.json", &incorrectStruct)
	if !errors.Is(err, errIncorrectStruct) {
		t.Errorf("incorrect struct not detected: %v", err)
	} else {
		t.Logf("👍: %v", err)
	}
}

func TestCorrectJsonCorrectStruct(t *testing.T) {
	err := Read("testJsonCorrect.json", &correctStruct)
	if err != nil {
		t.Errorf("there should be no error but there was one: %v", err)
	} else {
		t.Logf("👍: data parsed correctly")
	}
}
