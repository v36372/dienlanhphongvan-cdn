package client

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
)

type ImageService struct {
	client *Client
}

type imageOption struct {
	Width int `url:"width"`
}

func (s ImageService) Compress(src string) (io.Reader, error) {
	file, err := os.Open(src)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return s.compress(file, src)
}

func (s ImageService) CompressFile(file io.Reader, src string) (io.Reader, error) {
	return s.compress(file, src)
}

func (s ImageService) compress(file io.Reader, name string) (io.Reader, error) {
	path := "/images/compress"

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// this step is very important
	fileWriter, err := bodyWriter.CreateFormFile("file", name)
	if err != nil {
		return nil, err
	}

	// iocopy
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		return nil, err
	}

	contentType := bodyWriter.FormDataContentType()
	if err := bodyWriter.Close(); err != nil {
		return nil, err
	}

	req, err := s.client.NewUploadRequest(path, bodyBuf, contentType)
	if err != nil {
		return nil, err
	}

	w := &bytes.Buffer{}
	if _, err = s.client.Do(req, w); err != nil {
		return nil, err
	}
	return w, nil
}

type CropOption imageOption

func (s ImageService) Crop(src string, opt CropOption) (io.Reader, error) {
	file, err := os.Open(src)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return s.crop(file, src, opt)
}

func (s ImageService) CropFile(file io.Reader, src string, opt CropOption) (io.Reader, error) {
	return s.crop(file, src, opt)
}

func (s ImageService) crop(file io.Reader, name string, opt CropOption) (io.Reader, error) {
	path := "/images/crop"

	path, err := AddOptions(path, opt)
	if err != nil {
		return nil, err
	}

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// this step is very important
	fileWriter, err := bodyWriter.CreateFormFile("file", name)
	if err != nil {
		return nil, err
	}

	// iocopy
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		return nil, err
	}

	contentType := bodyWriter.FormDataContentType()
	if err := bodyWriter.Close(); err != nil {
		return nil, err
	}

	req, err := s.client.NewUploadRequest(path, bodyBuf, contentType)
	if err != nil {
		return nil, err
	}

	w := &bytes.Buffer{}
	if _, err = s.client.Do(req, w); err != nil {
		return nil, err
	}
	return w, nil
}

type ResizeOption imageOption

func (s ImageService) Resize(src string, opt ResizeOption) (io.Reader, error) {
	file, err := os.Open(src)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return s.resize(file, src, opt)
}

func (s ImageService) ResizeFile(file io.Reader, src string, opt ResizeOption) (io.Reader, error) {
	return s.resize(file, src, opt)
}

func (s ImageService) resize(file io.Reader, name string, opt ResizeOption) (io.Reader, error) {
	path := "/images/resize"

	path, err := AddOptions(path, opt)
	if err != nil {
		return nil, err
	}

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// this step is very important
	fileWriter, err := bodyWriter.CreateFormFile("file", name)
	if err != nil {
		return nil, err
	}

	// iocopy
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		return nil, err
	}

	contentType := bodyWriter.FormDataContentType()
	if err := bodyWriter.Close(); err != nil {
		return nil, err
	}

	req, err := s.client.NewUploadRequest(path, bodyBuf, contentType)
	if err != nil {
		return nil, err
	}

	w := &bytes.Buffer{}
	if _, err = s.client.Do(req, w); err != nil {
		return nil, err
	}
	return w, nil
}
