package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

const (
	retryFilesName      = "retries.txt"
	deleteRetryFileName = "deletes.txt"
	entryDelimiter      = "->"
)

type ProcessEntry struct {
	sourcePath string
	targetPath string
}

func processRetryFiles() {
	processDeletes()
	processUpdates()
}

func processDeletes() {
	//  read file if present
	readFile, err := os.Open(deleteRetryFileName)
	defer readFile.Close()
	if err != nil {
		log.Printf("Failed to read file: %s\n", err)
		return
	}
	entries := make([]string, 0)
	// read lines
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		entries = append(entries, fileScanner.Text())
	}
	// perform delete
	for _, entry := range entries {
		log.Printf("Deleting file %s", entry)
		err := os.Remove(entry)
		if err != nil {
			log.Printf("Failed to delete source file %s. Error: %s\n", entry, err)
			continue
		}
	}
}

func processUpdates() {
	//  read file if present
	readFile, err := os.Open(retryFilesName)
	defer readFile.Close()
	if err != nil {
		log.Printf("Failed to read file: %s\n", err)
		return
	}
	entries := make([]ProcessEntry, 0)
	// read lines
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		entries = append(entries, lineToEntry(fileScanner.Text()))
	}
	// perform copy
	for _, entry := range entries {
		log.Printf("Copying file %s to %s\n", entry.sourcePath, entry.targetPath)
		source, err := os.Open(entry.sourcePath)
		if err != nil {
			log.Printf("Failed to open source file %s. Error: %s\n", entry.sourcePath, err)
			continue
		}
		defer source.Close()
		target, err := os.OpenFile(entry.targetPath, os.O_RDWR, 0666)
		if err != nil {
			log.Printf("Failed to open target file %s. Error: %s\n", entry.targetPath, err)
			continue
		}
		defer target.Close()
		_, err = io.Copy(target, source)
		if err != nil {
			log.Printf("Failed to writer target file %s. Error: %s\n", entry.targetPath, err)
			continue
		}
	}
}

func lineToEntry(line string) ProcessEntry {
	parts := strings.Split(line, entryDelimiter)
	return ProcessEntry{
		sourcePath: parts[0],
		targetPath: parts[1],
	}
}
