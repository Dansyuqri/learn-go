package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("testing lint")
}

func validateDir(dirPath string) (valid bool) {
	_, err := ioutil.ReadDir(dirPath)

	return err == nil
}
