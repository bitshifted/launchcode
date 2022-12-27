package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

const (
	retryFilesName      = "retries.txt"
	deleteRetryFileName = "deletes.txt"
	entryDelimiter      = "->"
	cleanupDelaySeconds = 5
	oldExtension        = ".old"
)

var launcherUpdated = false

type ProcessEntry struct {
	sourcePath string
	targetPath string
}

func processRetryFiles(launcherFilePath string) {
	processDeletes()
	processUpdates(launcherFilePath)
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

func processUpdates(launcherFilePath string) {
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
		if entry.targetPath == launcherFilePath {
			processLauncherUpdate(entry)
		}
		log.Printf("Copying file %s to %s\n", entry.sourcePath, entry.targetPath)
		source, err := os.Open(entry.sourcePath)
		if err != nil {
			log.Printf("Failed to open source file %s. Error: %s\n", entry.sourcePath, err)
			continue
		}
		defer source.Close()
		target, err := os.OpenFile(entry.targetPath, os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			log.Printf("Failed to open target file %s. Error: %s\n", entry.targetPath, err)
			continue
		}
		defer target.Close()
		_, err = io.Copy(target, source)
		if err != nil {
			log.Printf("Failed to write target file %s. Error: %s\n", entry.targetPath, err)
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

func processLauncherUpdate(entry ProcessEntry) {
	log.Println("Processing launcher update")
	oldPath := entry.targetPath + oldExtension
	// check entry permissions
	info, err := os.Stat(entry.targetPath)
	if err != nil {
		log.Printf("Could not stat file %s. Error: %s\n", entry.targetPath, err)
	}
	info.Mode()
	os.Rename(entry.targetPath, oldPath)
	log.Printf("Current launcher renamed to %s\n", oldPath)
	launcherUpdated = true
	// create new file
	newFile, err := os.Create(entry.targetPath)
	if err != nil {
		log.Printf("Failed to create launcher copy: %s\n", err)
	}
	newFile.Chmod(info.Mode())
}

func cleanup(oldLaucnherFile string) {
	time.Sleep(cleanupDelaySeconds * time.Second)
	err := os.Remove(retryFilesName)
	if err != nil {
		log.Printf("Failed to delete file %s\n", retryFilesName)
	}
	err = os.Remove(deleteRetryFileName)
	if err != nil {
		log.Printf("Failed to delete file %s\n", deleteRetryFileName)
	}
	err = os.Remove(oldLaucnherFile)
	if err != nil {
		log.Printf("Failed to delete file %s\n", oldLaucnherFile)
	}
}
