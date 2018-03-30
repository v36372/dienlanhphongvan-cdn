package main

import (
	"dienlanhphongvan-cdn/client"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const (
	imgxAddress = "0.0.0.0:3000"
	srcImg      = "/Users/txc/Desktop/dau-da-xay-kungfu-tea-1496026661078433593.jpg"
)

func main() {
	ret, err := compress()
	if err != nil {
		fmt.Printf("ERROR | compress: %s, err: %v\n", srcImg, err)
	} else {
		fmt.Printf("OK    | compress: %s ==> %s\n", srcImg, absPath(ret))
	}
	ret, err = crop()
	if err != nil {
		fmt.Printf("ERROR | crop: %s, err: %v\n", srcImg, err)
	} else {
		fmt.Printf("OK    | crop: %s ==> %s\n", srcImg, absPath(ret))
	}
	ret, err = resize()
	if err != nil {
		fmt.Printf("ERROR | resize: %s, err: %v\n", srcImg, err)
	} else {
		fmt.Printf("OK    | resize: %s ==> %s\n", srcImg, absPath(ret))
	}
}

func compress() (string, error) {
	c := client.NewClient(fmt.Sprintf("http://%s", imgxAddress), nil)

	r, err := c.Image.Compress(srcImg)
	if err != nil {
		return "", err
	}
	path := "image_compress.jpg"
	if err := writeFile(path, r); err != nil {
		return "", err
	}
	return path, nil
}

func resize() (string, error) {
	c := client.NewClient(fmt.Sprintf("http://%s", imgxAddress), nil)

	r, err := c.Image.Resize(srcImg, client.ResizeOption{
		Width: 960,
	})
	if err != nil {
		return "", err
	}
	path := "image_resize.jpg"
	if err := writeFile(path, r); err != nil {
		return "", err
	}
	return path, nil
}

func crop() (string, error) {
	c := client.NewClient(fmt.Sprintf("http://%s", imgxAddress), nil)
	r, err := c.Image.Crop(srcImg, client.CropOption{
		Width: 960,
	})
	if err != nil {
		return "", err
	}
	path := "image_crop.jpg"
	if err := writeFile(path, r); err != nil {
		return "", err
	}
	return path, nil
}

func writeFile(path string, r io.Reader) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, r)
	if err != nil {
		return err
	}
	return nil
}

func absPath(file string) string {
	ret, err := filepath.Abs(file)
	if err != nil {
		return file
	}
	return ret

}
