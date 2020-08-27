package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func createTemp(exts []string) (tempDir string) {
	tempContent := []byte("tempContent")
	tempDir, err := ioutil.TempDir("", "tempDir")
	if err != nil {
		log.Fatal(err)
	}

	for _, ext := range exts {
		var strBuilder strings.Builder
		strBuilder.WriteString("tempfile")
		strBuilder.WriteString(ext)
		tmpFile := filepath.Join(tempDir, strBuilder.String())
		if err := ioutil.WriteFile(tmpFile, tempContent, 0666); err != nil {
			log.Fatal(err)
		}
	}

	return
}

func readDirFileNames(dirPath string) (fileNames []string) {
	fileInfo, err := ioutil.ReadDir(dirPath)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range fileInfo {
		fileNames = append(fileNames, file.Name())
	}
	return
}

func setupTestCase(exts []string, t *testing.T) (tempDir string, cleanUpFunc func(t *testing.T)) {
	tempDir = createTemp(exts)

	cleanUpFunc = func(t *testing.T) {
		os.RemoveAll(tempDir)
	}
	return
}

func TestCreateTemp(t *testing.T) {
	exts := []string{".py", ".txt", ".java"}
	tempFiles := []string{"tempfile.java", "tempfile.py", "tempfile.txt"}

	tempDir, cleanUpFunc := setupTestCase(exts, t)
	defer cleanUpFunc(t)

	fileNames := readDirFileNames(tempDir)
	if len(fileNames) != len(exts) {
		t.Errorf("want %v, got %v", len(exts), len(fileNames))
	}

	if !reflect.DeepEqual(fileNames, tempFiles) {
		t.Errorf("Files created: %v,  not equal to tempFiles: %v", fileNames, tempFiles)
	}
}

func TestGetUserInput(t *testing.T) {
	var stdin bytes.Buffer
	testStr := "testing123"

	stdin.Write([]byte(testStr))
	userInput := getUserInput(&stdin)

	if userInput != testStr {
		t.Errorf("Expected %v, Actual %v", testStr, userInput)
	}
}

func TestValidateDir(t *testing.T) {
	t.Run("Invalid Path", func(t *testing.T) {
		if validateDir("!@#$%^&**^%$#@!") {
			t.Error("Invalid path produced `true`")
		}
	})

	exts := []string{}

	tempDir, cleanUpFunc := setupTestCase(exts, t)
	defer cleanUpFunc(t)

	t.Run("Valid Path", func(t *testing.T) {
		if !validateDir(tempDir) {
			t.Error("Valid path produced `false`")
		}
	})
}
