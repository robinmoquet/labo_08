package zip

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Zipper it use for zip a folder on the output file
func Zipper(folderPath string, outputFile string) error {

	outputPathArray := strings.Split(outputFile, "/")
	outputFileName := outputPathArray[len(outputPathArray)-1]
	if len(outputPathArray) > 1 {
		outputPathArray = outputPathArray[:len(outputPathArray)-1]
		createOutputPathIfNotExist(outputPathArray)
	}

	if err := checkExtOfOutputFile(outputFileName); err != nil {
		return err
	}
	if err := checkIfFileNotExiste(outputFile); err != nil {
		return err
	}

	newZipFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("hear")
		return err
	}
	defer newZipFile.Close()

	return nil
}

// Check if the output file have the good ext
// Ex : something.zip
func checkExtOfOutputFile(fileName string) error {

	fileNameArray := strings.Split(fileName, ".")
	ext := fileNameArray[len(fileNameArray)-1]

	if ext != "zip" {
		return errors.New("You must use a valid extension ! Ex : something.zip")
	}

	return nil
}

// Check if the output file not existe
func checkIfFileNotExiste(outputFile string) error {

	if _, err := os.Stat(outputFile); err == nil {
		return errors.New("The file \"" + outputFile + "\" already exists")
	}
	return nil
}

// Create the output path if not exist
func createOutputPathIfNotExist(outputPathArray []string) {
	outputPath := filepath.Join(outputPathArray...)
	os.MkdirAll(outputPath, os.ModePerm)
}
