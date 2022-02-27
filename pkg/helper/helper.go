package helper

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/aditya109/library-system/internal/models"
	logger "github.com/sirupsen/logrus"
)

// GetAbsolutePath provides absolute path for a relative path -
func GetAbsolutePath(relPath string) (string, error) {
	var projectName = "library-system"
	if relPath[0] == '/' {
		relPath = relPath[1:]
	}
	splitRelativePath := strings.Split(relPath, "/")
	cwd, err := os.Getwd()
	var path string
	if err != nil {
		logger.Error(err)
		return "", err
	}
	projectLocation := strings.Split(cwd, projectName)
	path = filepath.Join(projectLocation[0], projectName, splitRelativePath[0], splitRelativePath[1])
	return path, nil
}

// GetFormattedFileName gets a formatted filename
func GetFormattedFileName(location models.PersistenceLocation) (string, error) {
	var timeStamp = time.Now().Format("2006-01-02_15:04:05")
	path := location.ContainerDirectory
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			logger.Error(err)
			return "", err
		}
	}
	var logFileName = fmt.Sprintf("%s/%s_%s.%s", location.ContainerDirectory, location.TargetFileName[0], timeStamp, location.TargetFileExtension)
	return logFileName, nil
}
