package main

import (
	"archive/zip"
	"io"
	"os"
	"sync"
)

func sequenced(files []string) {
	for _, fileName := range files {
		archive(fileName)
	}
}

func concurrent(files []string) {
	waitGroup := sync.WaitGroup{}
	for _, fileName := range files {
		fName := fileName
		waitGroup.Add(1)
		go func(fileName string) {
			defer func() {
				waitGroup.Done()
			}()
			archive(fName)
		}(fileName)
	}
	waitGroup.Wait()
}

func archive(fileName string) {
	archivedFiles, err := os.Create(archiveDirectory + fileName + zipFormat)
	if err != nil {
		return
	}

	defer func() {
		err = archivedFiles.Close()
		if err != nil {
			return
		}
	}()

	writer := zip.NewWriter(archivedFiles)

	defer func() {
		err = writer.Close()
		if err != nil {
			return
		}
	}()

	file, err := os.Open(filesDirectory + fileName)
	if err != nil {
		return
	}

	archive, err := writer.Create(fileName)
	if err != nil {
		return
	}
	_, err = io.Copy(archive, file)
	if err != nil {
		return
	}
}