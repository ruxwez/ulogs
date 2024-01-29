package ulogs

import (
	"errors"
	"fmt"
	"os"
	"time"
)

var (
	errToCreateFile = errors.New("error to create the file")
	errToWriteFile  = errors.New("error to write in the file")
)

type filesRepository struct {
	config *Config
}

// Function to record the logs in the logs file.
func (s *filesRepository) addLogFile(typeLog string, message string) error {
	var fileName string

	// Check if the file prefix is set
	if s.config.FilePrefix != "" {
		fileName = fmt.Sprintf("%s-%s.log", s.config.FilePrefix, time.Now().Format("2006-01-02"))
	} else {
		fileName = time.Now().Format("2006-01-02") + ".log"
	}

	// Open the file
	f, err := os.OpenFile("./"+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return errToCreateFile
	}

	// Close the file when this function is finished
	defer f.Close()

	// Write in the file the log message
	message = fmt.Sprintf("%s -> [%s] %s \n", time.Now().Format("2006-01-02 15:04:05"), typeLog, message)
	_, err = f.Write([]byte(message))

	if err != nil {
		return errToWriteFile
	}

	return nil
}
