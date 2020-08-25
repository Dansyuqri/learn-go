package main_test

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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
		strBuilder.WriteString("tmpfile")
		strBuilder.WriteString(ext)
		tmpFile := filepath.Join(tempDir, strBuilder.String())
		if err := ioutil.WriteFile(tmpFile, tempContent, 0666); err != nil {
			log.Fatal(err)
		}
	}

	return tempDir
}

func TestCreateTemp(t *testing.T) {
	exts := []string{".py", ".txt", ".java"}

	tempDir := createTemp(exts)
	defer os.RemoveAll(tempDir)
	fileInfo, err := ioutil.ReadDir(tempDir)

	if err != nil {
		log.Fatal(err)
	}

	if len(fileInfo) != len(exts) {
		t.Errorf("want %v, got %v", len(exts), len(fileInfo))
	}
}
