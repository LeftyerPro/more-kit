/*
 * kitComp.go Comp-difine
 * by:Leftyer
 * dt:2025-11-01
 */
package kitComp

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/chai2010/webp"
)

/* Comp-config-difine by:Leftyer dt:2025-11-01 */
type CompConfig struct {
	InputPath  string
	OutPutPath string
	FileType   int
}

/* Comp-Image by:Leftyer dt:2025-11-01 */
func (c *CompConfig) CompImg() error {
	f, err := os.Open(c.InputPath)
	if err != nil {
		return err
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		return err
	}
	o, err := os.Create(c.OutPutPath)
	if err != nil {
		return err
	}
	defer o.Close()
	switch c.FileType {
	case 1: /*PNG*/
		return (&png.Encoder{CompressionLevel: png.BestCompression}).Encode(o, img)
	case 2: /*JPEG*/
		return jpeg.Encode(o, img, &jpeg.Options{Quality: 100})
	case 3: /*WebP*/
		return webp.Encode(o, img, &webp.Options{Lossless: true, Quality: 100})
	}
	return nil
}
