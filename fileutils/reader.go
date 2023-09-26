package fileutils

import "os"

func ReadTextfile(fileName string) (string, error) {
	content, err := os.ReadFile(fileName)

	if err != nil {
		return "", err
	} else {
		return string(content), nil
	}
}
