package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("testing lint")
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

func safeRM(filePath string, stdin io.Reader) (pass bool) {
	fmt.Printf("Enter file path to confirm deletion: %v", filePath)
	userInput := getUserInput(stdin)
	_, err := os.Stat(filePath)
	if strings.TrimRight(userInput, "/n") == filePath && err == nil {
		os.Remove(filePath)
		return true
	}
	return false
}

func validateDir(dirPath string) (valid bool) {
	_, err := ioutil.ReadDir(dirPath)

	return err == nil
}
