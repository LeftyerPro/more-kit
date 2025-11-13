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
	if err := kitComp.CompImg(kitComp.CompConfig{InputPath: testPNG, OutPutPath: testOut, FileType: fileType}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(testOut); os.IsNotExist(err) {
		t.Fatal("output missing")
	}
	//os.Remove(testOut)
}

/* TestBenchmark-CompImg by:Leftyer dt:2025-11-01 */
func BenchmarkCompImg(b *testing.B) {
	for b.Loop() {
		_ = kitComp.CompImg(kitComp.CompConfig{InputPath: testPNG, OutPutPath: testOut, FileType: fileType})
		os.Remove(testOut)
	}
}
