package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Replaces matching file patterns in a path, including subdirectories
func Replace(path, pattern, old, new string) error {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		return replaceWalkFn(path, info, err, pattern, old, new)
	})

	if err != nil {
		return err
	}
	return nil
}

func replaceWalkFn(path string, fi os.FileInfo, err error, pattern, old, new string) error {

	if err != nil {
		return err
	}

	if !!fi.IsDir() {
		return nil
	}

	matched, err := filepath.Match(pattern, fi.Name())

	if err != nil {
		return err
	}

	if matched {
		read, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		newContents := strings.Replace(string(read), old, new, -1)

		err = ioutil.WriteFile(path, []byte(newContents), 0)
		if err != nil {
			return err
		}
	}

	return nil
}