package repo

import (
	"bytes"
	"errors"
	"io"
	"os"
	"testing"
)

func TestFileWriter(t *testing.T) {
	tests := []struct {
		name          string
		filename      string
		content       string
		expectedError error
	}{
		{
			name:          "Valid content",
			filename:      "testfile1.txt",
			content:       "This is a test.",
			expectedError: nil,
		},
		{
			name:          "Empty content",
			filename:      "testfile2.txt",
			content:       "",
			expectedError: nil,
		},
		{
			name:          "Invalid filename",
			filename:      "/nonexistent-path/testfile3.txt",
			content:       "This is a test.",
			expectedError: errors.New("Error writing to file"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &DataFileWriter{
				Filename: tt.filename,
			}

			err := writer.Write(tt.content)

			if tt.expectedError != nil && err == nil {
				t.Errorf("Expected error, but got none")
			}

			if tt.expectedError == nil && err != nil {
				t.Errorf("Error writing to file: %v", err)
			}

			if tt.expectedError == nil {
				file, err := os.Open(tt.filename)
				if err != nil {
					t.Errorf("Error opening test file: %v", err)
				}
				defer file.Close()

				var buf bytes.Buffer
				_, err = io.Copy(&buf, file)
				if err != nil {
					t.Errorf("Error reading test file: %v", err)
				}

				data := buf.String()

				if data != tt.content {
					t.Errorf("Expected content: %s, got: %s", tt.content, data)
				}

				err = os.Remove(tt.filename)
				if err != nil {
					t.Errorf("Error removing test file: %v", err)
				}
			}
		})
	}
}
