package cmd

import (
	"flag"

	_ "fmt"
	"log"
	"os"
	_ "strings"
)


const (
	zipFormat        = ".zip"
	filesDirectory   = "./files/"
	archiveDirectory = "./archive/"
	conn ="con"
	seqq="seq"
	defaultt ="default"

	)

var con = flag.String("con","default","con")
var seq = flag.String("seq","default","seq")

func main() {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("Can't close file; %v", err)
		}
	}()
	log.SetOutput(file)
	var cmd string
	flag.Parse()
	files := flag.Args()
	if *con != defaultt{
		cmd = conn

	}else if *seq != defaultt{
		cmd = seqq

	}
	operations(cmd,files)

}

func operations(cmd string, files []string) {

	switch cmd {
	case conn:
		concurrent(files)
	case seqq:
		sequenced(files)
	default:
		log.Printf("Not such flag!")
	}

}