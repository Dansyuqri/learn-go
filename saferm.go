package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var dirPathVar string

func main() {
	flag.Parse()
	if dirPathVar == "." {
		dirPathVar, _ = os.Getwd()
	}
	filenames := readDirFileNames(dirPathVar)
	for _, filename := range filenames {
		safeRM(dirPathVar, filename, os.Stdin)
	}
}

func init() {
	flag.StringVar(&dirPathVar, "p", ".", "Specify the path to run `saferm` on.")
}

func getUserInput(stdin io.Reader) string {
	reader := bufio.NewReader(stdin)
	userInput, _ := reader.ReadString('\n')
	return userInput
}

func readDirFileNames(dirPath string) (fileNames []string) {
	fileInfo, err := ioutil.ReadDir(dirPath)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range fileInfo {
		if !file.IsDir() {
			fileNames = append(fileNames, file.Name())
		}
	}
	return
}

func safeRM(dirPath string, filename string, stdin io.Reader) (pass bool) {
	fmt.Printf("Enter filename to confirm deletion: \n%v\n", filename)
	userInput := getUserInput(stdin)
	fullpath := filepath.Join(dirPath, filename)
	_, err := os.Stat(fullpath)
	if strings.TrimRight(userInput, "/n") == filename && err == nil {
		os.Remove(fullpath)
		return true
	}
	return false
}

func validateDir(dirPath string) (valid bool) {
	_, err := ioutil.ReadDir(dirPath)

	return err == nil
}
