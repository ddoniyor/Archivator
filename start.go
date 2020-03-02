package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	zipFormat        = ".zip"
	filesDirectory   = "./files/"
	archiveDirectory = "./archive/"
)


func main() {
	files := os.Args[1:]
	if files == nil {
		return
	}
	var cmd string

	fmt.Print("type command start to archive: ")
	_, err := fmt.Scan(&cmd)
	if err != nil {
		log.Printf("cant get input: %v", err)
		return
	}
	switch strings.ToLower(strings.TrimSpace(cmd)) {
	case "start":
		concurrent(files)
		sequenced(files)
		fmt.Printf("files succesfully archived!")
	}
	
}
