package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

// FileManager stores the paths used for reading input data
// and writing processed results to an output file.
type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

// ReadLines opens the configured input file and returns its contents
// as a slice of strings, where each element represents one line.
//
// The file is closed automatically with defer when this method finishes,
// whether it returns successfully or exits early because of an error.
func (fm FileManager) ReadLines() ([]string, error) {
	// Open the input file for reading.
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		return nil, errors.New("failed to open file")
	}

	// defer delays file.Close() until ReadLines finishes.
	//
	// This is safer than manually calling file.Close() in every possible
	// return path because Go guarantees that the deferred call runs when
	// the surrounding function exits.
	defer file.Close()

	// Create a scanner that reads the file one line at a time.
	scanner := bufio.NewScanner(file)

	var lines []string

	// Continue reading lines until there are no more lines to scan.
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// scanner.Err() reports any error that occurred while scanning.
	err = scanner.Err()

	if err != nil {
		return nil, errors.New("failed to read line in file")
	}

	return lines, nil
}

// WriteResult creates the configured output file and writes data into it
// as JSON.
//
// The output file is automatically closed with defer when this method ends,
// including cases where JSON encoding fails and the function returns early.
func (fm FileManager) WriteResult(data interface{}) error {
	// Create the output file.
	//
	// os.Create creates a new file or truncates the existing file
	// if a file with the same name already exists.
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("failed to create file")
	}

	// Ensure the file is closed when WriteResult finishes.
	//
	// Because this is deferred, we do not need to manually call
	// file.Close() before every return statement.
	defer file.Close()

	// Simulate a slow file-writing operation.
	// This is included for the course to make concurrent behavior easier to see.
	time.Sleep(3 * time.Second)

	// Create a JSON encoder that writes directly into the file.
	encoder := json.NewEncoder(file)

	// Convert data to JSON and write it into the output file.
	err = encoder.Encode(data)

	if err != nil {
		return errors.New("failed to convert data to JSON")
	}

	return nil
}

// New creates and returns a FileManager configured with
// the input file path and output file path provided.
func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
