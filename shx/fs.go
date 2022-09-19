package shx

import (
	"os"
	"path/filepath"
)

func GetWorkingDirectory() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Abs(pwd)
}
