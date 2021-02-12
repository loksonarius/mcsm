package integration

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type ErrorList []error

func (e ErrorList) Error() string {
	var out strings.Builder

	for _, err := range e {
		out.WriteString(err.Error())
		out.WriteRune('\n')
	}

	return out.String()
}

func copyDir(source, destination string) error {
	walkFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relativePath := strings.Replace(path, source, "", 1)
		if relativePath == "" {
			return nil
		}

		targetPath := filepath.Join(destination, relativePath)
		if info.IsDir() {
			return os.Mkdir(targetPath, info.Mode())
		} else if !info.Mode().IsRegular() {
			return fmt.Errorf("unexpected non-regular file")
		}

		srcFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		dstFile, err := os.Create(targetPath)
		if err != nil {
			return err
		}
		defer dstFile.Close()

		if _, err := io.Copy(dstFile, srcFile); err != nil {
			return err
		}

		if err := dstFile.Chmod(info.Mode()); err != nil {
			return err
		}

		return nil
	}

	return filepath.Walk(source, walkFunc)
}

func withSuiteDir(suite string, f func() error) error {
	serverDir, err := ioutil.TempDir("servers", suite+"-*")
	if err != nil {
		return fmt.Errorf("failed to setup server dir for %s: %s", suite, err)
	}

	serverDir, err = filepath.Abs(serverDir)
	if err != nil {
		return fmt.Errorf("failed to get abs path for %s", serverDir)
	}
	deleteServerDir := true
	defer func() {
		if deleteServerDir {
			os.RemoveAll(serverDir)
		}
	}()

	suiteDir := filepath.Join("suites", suite+".suite")
	copyDir(suiteDir, serverDir)

	workingDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("unable to get current working directory: %s", err)
	}

	if err = os.Chdir(serverDir); err != nil {
		deleteServerDir = false
		return fmt.Errorf("failed to chdir into %s", serverDir)
	}
	defer os.Chdir(workingDir)

	err = f()
	if err != nil {
		deleteServerDir = false
	}

	return err
}

func asTest(t *testing.T, f func() error) {
	if err := f(); err != nil {
		t.Error(err)
	}
}
