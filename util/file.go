package util

import (
	"io"
	"io/ioutil"
	"os"
	"path"
)

func WithTempFile(filepath string, handle func(tmpfile *os.File) error) error {
	tmpfile, err := ioutil.TempFile("", path.Base(filepath))
	if err != nil {
		return err
	}
	defer os.Remove(tmpfile.Name())
	if err := handle(tmpfile); err != nil {
		return err
	}
	return nil
}

func ExistFile(path string) bool {
	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	if stat.IsDir() {
		return false
	}
	return true
}

func WriteFile(filepath string, r io.Reader) error {
	if err := os.MkdirAll(path.Dir(filepath), os.ModePerm); err != nil {
		return err
	}
	w, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			os.Remove(filepath)
		}
	}()
	_, err = io.Copy(w, r)
	return err
}

func MoveFile(srcFilepath, dstFilepath string) error {
	if err := os.MkdirAll(path.Dir(dstFilepath), os.ModePerm); err != nil {
		return err
	}
	return os.Rename(srcFilepath, dstFilepath)
}
