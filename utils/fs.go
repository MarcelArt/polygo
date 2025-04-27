package utils

import "os"

func PWD() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return pwd, nil
}

func RM(path string) error {
	return os.RemoveAll(path)
}
