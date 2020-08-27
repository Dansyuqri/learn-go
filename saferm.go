package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	fmt.Println("testing lint")
}

func getUserInput(stdin io.Reader) string {
	fmt.Print("Enter text: ")
	reader := bufio.NewReader(stdin)
	userInput, _ := reader.ReadString('\n')
	return userInput
}

func validateDir(dirPath string) (valid bool) {
	_, err := ioutil.ReadDir(dirPath)

	return err == nil
}
