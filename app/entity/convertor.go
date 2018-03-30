package entity

import (
	"bytes"
	"dienlanhphongvan-cdn/util"
	"fmt"
	"io"
	"os"
	"os/exec"
)

type Convertor struct {
	enable bool
	exec   string
}

func NewConvertor(enable bool, execPath string) *Convertor {
	return &Convertor{
		enable: enable,
		exec:   execPath,
	}
}

func (c Convertor) Enable() bool {
	return c.enable
}

func (c Convertor) Convert(r io.Reader) (io.Reader, error) {
	if !c.enable {
		return nil, fmt.Errorf("convert: unavailable, current is disable")
	}
	var (
		buf = &bytes.Buffer{}
		err error
	)
	err = util.WithTempFile("", func(f *os.File) error {
		if _, err := io.Copy(f, r); err != nil {
			return err
		}
		if err := f.Close(); err != nil {
			return err
		}
		path := f.Name()
		errBuf := &bytes.Buffer{}
		cmd := exec.Command(c.exec, path, path)
		cmd.Stderr = errBuf
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("convert: %s %s", err.Error(), errBuf.String())
		}
		f, err = os.Open(f.Name())
		if err != nil {
			return err
		}
		defer f.Close()
		if _, err := io.Copy(buf, f); err != nil {
			return err
		}
		return nil
	})
	return buf, err
}
