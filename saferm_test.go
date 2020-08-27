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

func compareDirFiles(tmpFile string, tempDir string, stdin bytes.Buffer, expectedRMFlag bool, altered []string, t *testing.T) {
	safeRMFlag := safeRM(tmpFile, &stdin)
	safeRMFiles := readDirFileNames(tempDir)
	if safeRMFlag != expectedRMFlag || !reflect.DeepEqual(altered, safeRMFiles) {
		t.Errorf("Files safeRM: %v,  not equal to altered: %v", safeRMFiles, altered)
	}
}

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

func TestReadDirFileNames(t *testing.T) {
	exts := []string{".py", ".txt", ".java"}
	tempFiles := []string{"tempfile.java", "tempfile.py", "tempfile.txt"}

	tempDir, cleanUpFunc := setupTestCase(exts, t)
	defer cleanUpFunc(t)

	fileNames := readDirFileNames(tempDir)

	if !reflect.DeepEqual(fileNames, tempFiles) {
		t.Errorf("Files created: %v,  not equal to tempFiles: %v", fileNames, tempFiles)
	}
}

func TestSafeRM(t *testing.T) {
	exts := []string{".py", ".txt", ".jpg", ".java"}
	altered := []string{"tempfile.java", "tempfile.jpg", "tempfile.txt"}
	tempDir, cleanUpFunc := setupTestCase(exts, t)
	defer cleanUpFunc(t)

	t.Run("Delete tempfile.py", func(t *testing.T) {
		var stdin bytes.Buffer
		tmpFile := filepath.Join(tempDir, "tempfile.py")

		stdin.Write([]byte(tmpFile))
		compareDirFiles(tmpFile, tempDir, stdin, true, altered, t)
	})

	t.Run("Delete non-existent file", func(t *testing.T) {
		var stdin bytes.Buffer
		tmpFile := filepath.Join(tempDir, "abc123.xyz")

		stdin.Write([]byte(tmpFile))

		compareDirFiles(tmpFile, tempDir, stdin, false, altered, t)
	})

}
