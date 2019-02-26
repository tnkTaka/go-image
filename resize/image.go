package resize

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/nfnt/resize"
)

const (
	IMAGE_PATH = "./xxx.xxx"
)

func ResizeImage() {
	file, err := os.Open(IMAGE_PATH)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, format, err := image.Decode(file) // resp.Body
	if err != nil {
		log.Fatal(err)
	}

	// 変換ロジック
	img = resize.Resize(500, 500, img, resize.Lanczos3)
	buf := new(bytes.Buffer)

	switch format {
	case "jpeg", "jpg":
		err = jpeg.Encode(buf, img, nil)
	case "png":
		err = png.Encode(buf, img)
	case "gif":
		err = gif.Encode(buf, img, nil)
	default:
		err = errors.New("invalid format")
	}

	if err != nil {
		log.Fatal(err)
	}

	// base64 encode
	bit := buf.Bytes()
	imgBase64 := base64.StdEncoding.EncodeToString([]byte(bit))

	// base64 decode
	data, _ := base64.StdEncoding.DecodeString(imgBase64)
	out, err := os.Create(fmt.Sprintf("../after.%s", format))
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	out.Write(data)
}
