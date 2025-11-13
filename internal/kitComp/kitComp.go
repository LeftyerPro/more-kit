/*
 * kitComp.go Comp-difine
 * by:Leftyer
 * dt:2025-11-01
 */
package kitComp

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/chai2010/webp"
)

/* Comp-config-difine by:Leftyer dt:2025-11-01 */
type CompConfig struct {
	InputPath  string
	OutPutPath string
	FileType   int
	Lossless   bool
	Quality    float32
}

/* Comp-Image by:Leftyer dt:2025-11-01 */
func CompImg(c CompConfig) error {
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
	return fmt.Errorf("unknown FileType %d", c.FileType)
}

/* Comp-Any2WebP by:Leftyer dt:2025-11-01 */
/* 输入：PDF / 图片（jpg/png/gif/tiff/bmp）*/
/* 输出：单文件 WebP（无损或质量可配）*/
func CompWebpBr(c CompConfig) error {
	/* 按扩展名分流 */
	ext := strings.ToLower(filepath.Ext(c.InputPath))
	switch ext {
	case ".pdf":
		return pdf2WebP(c)
	case ".gif":
		return gif2WebP(c)
	case ".png", ".jpg", ".jpeg", ".tiff", ".bmp", ".tif":
		return staticImg2WebP(c)
	}
	return fmt.Errorf("unsupported input ext: %s", ext)
}

/* pdf2WebP by:Leftyer dt:2025-11-01 */
/* 仅处理第 1 页；需要 go-fitz */
func pdf2WebP(c CompConfig) error {
	// doc, err := fitz.New(c.InputPath)
	// if err != nil {
	// 	return err
	// }
	// defer doc.Close()
	// if doc.NumPage() < 1 {
	// 	return fmt.Errorf("empty pdf")
	// }
	// img, err := doc.Image(0, 144) /* 默认 144 dpi */
	// if err != nil {
	// 	return err
	// }
	return saveWebP(c.OutPutPath, nil, c.Lossless, c.Quality)
}

/* gif2WebP by:Leftyer dt:2025-11-01 */
/* 取 GIF 第一帧 */
func gif2WebP(c CompConfig) error {
	f, err := os.Open(c.InputPath)
	if err != nil {
		return err
	}
	defer f.Close()
	first, err := gif.Decode(f)
	if err != nil {
		return err
	}
	return saveWebP(c.OutPutPath, first, c.Lossless, c.Quality)
}

/* staticImg2WebP by:Leftyer dt:2025-11-01 */
/* 通用单帧图片转 WebP */
func staticImg2WebP(c CompConfig) error {
	f, err := os.Open(c.InputPath)
	if err != nil {
		return err
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		return err
	}
	return saveWebP(c.OutPutPath, img, c.Lossless, c.Quality)
}

/* saveWebP by:Leftyer dt:2025-11-01 */
/* 统一写 WebP 文件 */
func saveWebP(outPath string, img image.Image, lossless bool, quality float32) error {
	o, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer o.Close()
	return webp.Encode(o, img, &webp.Options{Lossless: lossless, Quality: quality})
}
