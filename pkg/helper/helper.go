package helper

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/aditya109/library-system/internal/models"
	logger "github.com/sirupsen/logrus"
)

// GetAbsolutePath provides absolute path for a relative path -
func GetAbsolutePath(relPath string) string {
	if relPath[0] == '/' {
		relPath = relPath[1:]
	}
	splitRelativePath := strings.Split(relPath, "/")
	cwd, err := os.Getwd()
	var path string
	if err != nil {
		log.Fatal(err)
	}
	projectLocation := strings.Split(cwd, "go-server-template")
	path = filepath.Join(projectLocation[0], "go-server-template", splitRelativePath[0], splitRelativePath[1])
	return path
}

func GetFormattedFileName(specs models.PersistenceLocation) (string, error) {
	var timeStamp = time.Now().Format("2006-01-02_15:04:05")
	path := specs.ContainerDirectory
	if _, err := os.Stat(path); error.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			logger.Error(err)
			return "", err
		}
	}
}
