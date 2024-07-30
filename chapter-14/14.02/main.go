package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

var (
	ErrWorkingFileNotFound = errors.New("the working file is not found")
)

func main() {
	backupFilename := "backupFile.txt"
	workingFilename := "notes.txt"
	data := "note"
	if err := createBackupFile(workingFilename, backupFilename); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := 1; i <= 10; i++ {
		note := data + " " + strconv.Itoa(i)
		if err := addNotes(workingFilename, note); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	fmt.Println("Finished writing new notes.")
}

func createBackupFile(workingFilename, backupfileName string) error {
	_, err := os.Stat(workingFilename)
	if err != nil {
		// checking if the working file actually exists and throw a file not found error if it doesn't
		if os.IsNotExist(err) {
			return ErrWorkingFileNotFound
		}
		return err
	}
	// opening a connection to the workfile
	workingFile, err := os.Open(workingFilename)
	if err != nil {
		return err
	}
	/*
		The osFile implements the io.Reader interface so it can be passed to io.ReadAll.
		The content will be a byte of the entire file so this is not great for large files
		as it would take up a lot of memory.
	*/
	content, err := io.ReadAll(workingFile)
	fmt.Println(content)
	if err != nil {
		return err
	}

	err = os.WriteFile(backupfileName, content, 0644)
	if err != nil {
		return err
	}
	// returning nil at the end, indicating no errors occurred during execution of the function
	return nil
}

func addNotes(workingFilename, notes string) error {
	notes += "\n"

	// will create file if it doesn't exist
	file, err := os.OpenFile(
		workingFilename,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		return err
	}
	// closing the file connection once all code is done running within function
	defer file.Close()
	// write the added notes to the file.
	if _, err := file.Write([]byte(notes)); err != nil {
		return nil
	}

	return nil
}
