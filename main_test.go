package main_test

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func createTemp() (tempDir string) {
	tempContent := []byte("tempContent")
	tempDir, err := ioutil.TempDir("", "tempDir")
	if err != nil {
		log.Fatal(err)
	}

	tmpFile := filepath.Join(tempDir, "tmpfile")
	if err := ioutil.WriteFile(tmpFile, tempContent, 0666); err != nil {
		log.Fatal(err)
	}

	return tempDir
}

func TestCreateTemp(t *testing.T) {
	tempDir := createTemp()
	defer os.RemoveAll(tempDir)
	fileInfo, err := ioutil.ReadDir(tempDir)

	if err != nil {
		log.Fatal(err)
	}

	if len(fileInfo) != 1 {
		t.Errorf("want 1, got %v", len(fileInfo))
	}
}
