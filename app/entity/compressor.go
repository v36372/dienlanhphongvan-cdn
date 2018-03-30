package entity

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

type Compressor struct {
	enable bool
	exec   string
}

func NewCompress(enable bool, execPath string) *Compressor {
	return &Compressor{
		enable: enable,
		exec:   execPath,
	}
}

func (c Compressor) Enable() bool {
	return c.enable
}

func (c Compressor) Compress(r io.Reader) (io.Reader, error) {
	if !c.enable {
		return nil, fmt.Errorf("cjepg: unavailable, current is disable")
	}
	var (
		buf    bytes.Buffer
		errBuf bytes.Buffer
	)

	cmd := exec.Command(c.exec)
	cmd.Stdin = r
	cmd.Stdout = &buf
	cmd.Stderr = &errBuf
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("cjpeg: %s %s", err.Error(), errBuf.String())
	}
	return &buf, nil
}
