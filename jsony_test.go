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
	err := jsony("nonexistingfile.json", &correctStruct)
	if !errors.Is(err, ErrReadingFile) {
		t.Errorf("inexisting file not detected: %v", err)
	} else {
		t.Logf("👍: %v", err)
	}
}

func TestIncorrectJson(t *testing.T) {
	err := jsony("testJsonIncorrect.json", &correctStruct)
	if !errors.Is(err, ErrIncorrectJson) {
		t.Errorf("incorrect JSON not detected: %v", err)
	} else {
		t.Logf("👍: %v", err)
	}
}

func TestCorrectJsonIncorrectStruct(t *testing.T) {
	err := jsony("testJsonCorrect.json", &incorrectStruct)
	if !errors.Is(err, ErrIncorrectStruct) {
		t.Errorf("incorrect struct not detected: %v", err)
	} else {
		t.Logf("👍: %v", err)
	}
}

func TestCorrectJsonCorrectStruct(t *testing.T) {
	err := jsony("testJsonCorrect.json", &correctStruct)
	if err != nil {
		t.Errorf("there should be no error but there was one: %v", err)
	} else {
		t.Logf("👍: data parsed correctly")
	}
}

