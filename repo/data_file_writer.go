package repo

import (
	"io"
	"log"
	"os"
)

type DataFileWriter struct {
	Filename string
}

func (w *DataFileWriter) Write(data string) error {
	file, err := os.OpenFile(w.Filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		log.Printf("Error writing to file: %v", err)
		return err
	}

	return nil
}
