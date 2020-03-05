package archivator

import (

	"archive/zip"
	"io"
	"log"
	"os"
	"sync"
)

const  (
	ZipFormat        = ".zip"
	FilesDirectory   = "./files/"
	ArchiveDirectory = "./archive/"
)

func Sequenced(files []string) {
	for _, fileName := range files {
		archive(fileName)
	}
}

func Concurrent(files []string) {
	waitGroup := sync.WaitGroup{}
	for _, fileName := range files {
		fName := fileName
		waitGroup.Add(1)
		go func() {
			defer func() {
				waitGroup.Done()
			}()
			archive(fName)
		}()
	}
	waitGroup.Wait()
}

func archive(fileName string) {
	archivedFiles, err := os.Create(ArchiveDirectory + fileName + ZipFormat)
	if err != nil {
		log.Printf("cant create file in directory!")
		return
	}
	defer func() {
		err = archivedFiles.Close()
		if err != nil {
			log.Printf("cant close archived files!")
			return
		}
	}()
	writer := zip.NewWriter(archivedFiles)
	defer func() {
		err = writer.Close()
		if err != nil {
			log.Printf("cant close the file!")
			return
		}
	}()

	file, err := os.Open(FilesDirectory + fileName)
	if err != nil {
		return
	}
	archive, err := writer.Create(fileName)
	if err != nil {
		log.Printf("cant create file!")
		return
	}
	_, err = io.Copy(archive, file)
	if err != nil {
		log.Printf("cant copy the file!")
		return
	}
}