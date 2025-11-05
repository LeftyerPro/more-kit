/*
 * morekit.go Kit-difine
 * by:Leftyer
 * dt:2025-11-01
 */
package morekit

import (
	"github.com/LeftyerPro/more-kit/internal/kitComp"
)

/* CompRun-Image by:Leftyer dt:2025-11-01 */
func CompImage(input, output string, fileType int) error {
	c := &kitComp.CompConfig{
		InputPath:  input,
		OutPutPath: output,
		FileType:   fileType,
	}
	return c.CompImg()
}
