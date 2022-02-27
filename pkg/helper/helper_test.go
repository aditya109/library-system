package helper

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGetAbsolutePathForCorrectRelativePath(t *testing.T) {
	result, err := GetAbsolutePath(`/config/config.json`)
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	var expectedResult string
	cwd, err := os.Getwd()
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	expectedResult = filepath.Join(strings.Split(cwd, "pkg/helper")[0], "config/config.json")
	if result != expectedResult {
		t.Errorf("got %q, wanted %q", result, expectedResult)
	}
}
