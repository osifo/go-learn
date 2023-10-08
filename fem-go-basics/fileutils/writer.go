package fileutils

import (
	"os"
)

func WriteToFile(fileName string, content string) error {
	workingDirPath, _ := os.Getwd()
	pathToWrite := string(workingDirPath + "/" + fileName)
	return os.WriteFile(pathToWrite, []byte(content), 0666)
}
