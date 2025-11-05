/*
 * morekit_Comp_test.go Test-difine
 * by:Leftyer
 * dt:2025-11-01
 */
package morekit

import (
	"os"
	"testing"

	"github.com/LeftyerPro/more-kit/internal/kitComp"
)

const (
	testPNG  = `F:\Test\input\1.png`
	testOut  = `F:\Test\output\1.png`
	fileType = 1
)

/* TestUnit-CompImg by:Leftyer dt:2025-11-01 */
func TestCompImg(t *testing.T) {
	c := &kitComp.CompConfig{InputPath: testPNG, OutPutPath: testOut, FileType: fileType}
	if err := c.CompImg(); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(testOut); os.IsNotExist(err) {
		t.Fatal("output missing")
	}
	os.Remove(testOut)
}

/* TestBenchmark-CompImg by:Leftyer dt:2025-11-01 */
func BenchmarkCompImg(b *testing.B) {
	c := &kitComp.CompConfig{InputPath: testPNG, OutPutPath: testOut, FileType: fileType}
	for b.Loop() {
		_ = c.CompImg()
		os.Remove(testOut)
	}
}
