package entity

import (
	"bytes"
	"dienlanhphongvan-cdn/errors"
	"io"
	"io/ioutil"

	bimg "gopkg.in/h2non/bimg.v1"
)

type ImageProcessFunc func(buf []byte, width int) ([]byte, error)

type Image struct {
	Convertor  *Convertor
	Compressor *Compressor
}

func NewImage(convertor *Convertor, compressor *Compressor) *Image {
	return &Image{
		Convertor:  convertor,
		Compressor: compressor,
	}
}

func (r Image) Compress(input io.Reader) (io.Reader, error) {
	ret, err := r.Compressor.Compress(input)
	if err != nil {
		return nil, errors.ErrorInternalServer(err)
	}
	return ret, nil
}

func (r Image) Crop(input io.ReadCloser, width int) (io.Reader, error) {
	return r.process(input, width, crop)
}

func (r Image) Resize(input io.ReadCloser, width int) (io.Reader, error) {
	return r.process(input, width, resize)
}

func (r Image) process(input io.Reader, width int, processFunc ImageProcessFunc) (io.Reader, error) {
	// read
	buf, err := ioutil.ReadAll(input)
	if err != nil {
		return nil, errors.ErrorInternalServer(err)
	}
	// process
	newBuf, err := processFunc(buf, width)
	if err != nil {
		// try fixing invalid image by cmd: convert
		newInput, err := r.Convertor.Convert(bytes.NewBuffer(buf))
		if err != nil {
			// try fixing invalid image by cmd: cjpeg
			newInput, err = r.Compressor.Compress(bytes.NewBuffer(buf))
		}
		if err != nil {
			return nil, errors.ErrorInternalServer(err)
		}
		// read
		buf, err = ioutil.ReadAll(newInput)
		if err != nil {
			return nil, errors.ErrorInternalServer(err)
		}
		// process
		newBuf, err = processFunc(buf, width)
		if err != nil {
			return nil, errors.ErrorInternalServer(err)
		}
	}
	return bytes.NewBuffer(newBuf), nil
}

func resize(buf []byte, width int) ([]byte, error) {
	return bimg.NewImage(buf).Resize(width, 0)
}

func crop(buf []byte, width int) ([]byte, error) {
	// zoom before crop
	image := bimg.NewImage(buf)
	meta, err := image.Metadata()
	if err != nil {
		return nil, err
	}
	var zoomFactor = 1
	if meta.Size.Width < width {
		zoomFactor = width / meta.Size.Width
		if width%meta.Size.Width != 0 {
			zoomFactor += 1
		}
	}
	if zoomFactor > 1 {
		buf, err = image.Zoom(zoomFactor)
		if err != nil {
			return nil, err
		}
	}
	// resize then crop center
	options := bimg.Options{
		Width:   width,
		Height:  width,
		Embed:   true,
		Crop:    true,
		Gravity: bimg.GravityCentre,
	}
	return bimg.NewImage(buf).Process(options)
}
