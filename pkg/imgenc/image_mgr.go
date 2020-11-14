package imgenc

import (
	"errors"
	"image"
	"io"
	"io/ioutil"
	"log"

	zstd "github.com/valyala/gozstd"

	/*
		image/jpeg - Support of JPEG format
		image/png  - Support of PNG format
	*/
	_ "image/jpeg"
	_ "image/png"
)

// CImage -- Structure to store an image with format and data
type CImage struct {
	Format string
	// JPEG or PNG file, not raw pixels
	Raw []byte
}

// CreateCImage -- do create an CImage structure and fills it
func CreateCImage(r io.Reader) *CImage {
	_, fmt, err := image.DecodeConfig(r)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	mImage := CImage{Format: fmt, Raw: nil}

	mImage.Raw, _ = ioutil.ReadAll(r)
	log.Printf("Found %v in stream with size of %d bytes\n", mImage.Format, len(mImage.Raw))

	return &mImage
}

// Compress -- Compresses CImage raw bytes (png, jpeg files)
func (c *CImage) Compress() ([]byte, error) {
	if c.Raw == nil {
		log.Fatal("c.Raw is nil")
		return nil, errors.New("c.Raw have no content to compress")
	}

	return zstd.Compress(nil, c.Raw), nil
}
