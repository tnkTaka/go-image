package main

import (
	"bytes"
	"encoding/base64"
	"image/jpeg"
	"log"
	"os"

	"github.com/nfnt/resize"
)

func main() {
	encodePath := "./one.png" // request image path
	decodePath := "./one.jpg" // Image path after resize

	imageData, err := imageBase64Encode(encodePath, 500, 500)
	if err != nil {
		// error
		log.Fatal(err)
	}

	err = imageBase64Decode(decodePath, imageData)
	if err != nil {
		// error
		log.Fatal(err)
	}
}

func imageBase64Encode(path string, width, height uint) (string, error) {
	// file open
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return "", err
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		return "", err
	}

	// image resize logic
	m := resize.Resize(width, height, img, resize.Bilinear)
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, m, nil)
	if err != nil {
		return "", err
	}

	// image base64 encode
	bit := buf.Bytes()
	imgBase64 := base64.StdEncoding.EncodeToString([]byte(bit))

	return imgBase64, nil
}

func imageBase64Decode(path, imgBase64 string) error {
	// image base64 decode
	data, _ := base64.StdEncoding.DecodeString(imgBase64)

	// create image file
	out, err := os.Create(path)
	defer out.Close()
	if err != nil {
		return err
	}
	out.Write(data)

	return nil
}
