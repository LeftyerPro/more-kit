/*
 * morekit_File_test.go Test-difine
 * by:Leftyer
 * dt:2025-11-01
 */
package morekit

import (
	"os"
	"path/filepath"
	"testing"
)

var tmp = filepath.Join(`F:\Test`, "morekit_test")

/* TestUnit-FolderIsExist by:Leftyer dt:2025-11-01 */
func TestFolderIsExist(t *testing.T) {
	if !FolderIsExist(tmp, 0) {
		t.Fatal("nonExist should return false")
	}
}

/* TestBenchmark-FolderIsExist by:Leftyer dt:2025-11-01 */
func BenchmarkFolderIsExist(b *testing.B) {
	_ = os.MkdirAll(tmp, 0755)
	defer os.RemoveAll(tmp)
	for b.Loop() {
		_ = FolderIsExist(tmp, 0)
	}
}

/* TestUnit-FolderCpoy by:Leftyer dt:2025-11-01 */
func TestFolderCpoy(t *testing.T) {
	src := filepath.Join(tmp, "input")
	dst := filepath.Join(tmp, "output")
	if err := FolderCpoy(src, dst); err != nil {
		t.Fatal(err)
	}
}

/* TestBenchmark-FolderCpoy by:Leftyer dt:2025-11-01 */
func BenchmarkFolderCpoy(b *testing.B) {
	src := filepath.Join(tmp, "benchSrc")
	dst := filepath.Join(tmp, "benchDst")
	_ = os.RemoveAll(src)
	_ = os.MkdirAll(filepath.Join(src, "sub"), 0755)
	_ = os.WriteFile(filepath.Join(src, "f.txt"), []byte("bench"), 0644)

	for i := 0; b.Loop(); i++ {
		_ = FolderCpoy(src, dst+string(rune(i)))
	}
}

/* TestUnit-FileIsExist by:Leftyer dt:2025-11-01 */
func TestFileIsExist(t *testing.T) {
	f := filepath.Join(tmp, "test.txt")
	_ = os.Remove(f)
	defer os.Remove(f)
	if FileIsExist(f) {
		t.Fatal("nonExist file should return false")
	}
	_ = os.WriteFile(f, []byte("x"), 0644)
	if !FileIsExist(f) {
		t.Fatal("exist file should return true")
	}
}

/* TestBenchmark-FileIsExist by:Leftyer dt:2025-11-01 */
func BenchmarkFileIsExist(b *testing.B) {
	f := filepath.Join(tmp, "benchExist.txt")
	_ = os.WriteFile(f, []byte("x"), 0644)
	defer os.Remove(f)
	for b.Loop() {
		_ = FileIsExist(f)
	}
}

/* TestUnit-FileCpoy by:Leftyer dt:2025-11-01 */
func TestFileCpoy(t *testing.T) {
	src := filepath.Join(tmp, "input", "1.png")
	dst := filepath.Join(tmp, "output", "1.png")
	if err := FileCpoy(src, dst); err != nil {
		t.Fatal(err)
	}
	if !FileIsExist(dst) {
		t.Fatal("dst file not copied")
	}
}

/* TestBenchmark-FileCpoy by:Leftyer dt:2025-11-01 */
func BenchmarkFileCpoy(b *testing.B) {
	src := filepath.Join(tmp, "benchSrc.txt")
	dst := filepath.Join(tmp, "benchDst.txt")
	_ = os.WriteFile(src, []byte("bench"), 0644)
	defer os.Remove(src)

	for i := 0; b.Loop(); i++ {
		_ = FileCpoy(src, dst+string(rune(i)))
	}
}

/* File-Read-Test by:Leftyer dt:2025-11-01 */
func TestFileRead(t *testing.T) {
	f := filepath.Join(tmp, "read.txt")
	_ = os.WriteFile(f, []byte("hello"), 0644)
	defer os.Remove(f)

	s, err := FileRead(f)
	if err != nil || s != "hello" {
		t.Fatalf("FileRead fail: %v", err)
	}
}

/* Benchmark-File-Read by:Leftyer dt:2025-11-01 */
func BenchmarkFileRead(b *testing.B) {
	f := filepath.Join(tmp, "benchRead.txt")
	_ = os.WriteFile(f, []byte("benchmark"), 0644)
	defer os.Remove(f)
	for b.Loop() {
		_, _ = FileRead(f)
	}
}

/* File-Write-Test by:Leftyer dt:2025-11-01 */
func TestFileWrite(t *testing.T) {
	f := filepath.Join(tmp, "write.txt")
	defer os.Remove(f)

	if err := FileWrite(f, "world"); err != nil {
		t.Fatalf("FileWrite fail: %v", err)
	}
	if b, _ := os.ReadFile(f); string(b) != "world" {
		t.Fatal("FileWrite content mismatch")
	}
}

/* Benchmark-File-Write by:Leftyer dt:2025-11-01 */
func BenchmarkFileWrite(b *testing.B) {
	f := filepath.Join(tmp, "benchWrite.txt")
	defer os.Remove(f)
	for b.Loop() {
		_ = FileWrite(f, "benchmark")
	}
}

/* Json-Read-Test by:Leftyer dt:2025-11-01 */
func TestJsonRead(t *testing.T) {
	f := filepath.Join(tmp, "read.json")
	_ = os.WriteFile(f, []byte(`{"Name":"Leftyer","Age":18}`), 0644)
	defer os.Remove(f)

	type U struct {
		Name string
		Age  int
	}
	v, err := JsonRead[U](f)
	if err != nil || v.Name != "Leftyer" || v.Age != 18 {
		t.Fatalf("JsonRead fail: %v", err)
	}
}

/* Benchmark-Json-Read by:Leftyer dt:2025-11-01 */
func BenchmarkJsonRead(b *testing.B) {
	f := filepath.Join(tmp, "benchRead.json")
	_ = os.WriteFile(f, []byte(`{"Name":"bench","Age":99}`), 0644)
	defer os.Remove(f)
	type U struct {
		Name string
		Age  int
	}
	for b.Loop() {
		_, _ = JsonRead[U](f)
	}
}

/* Json-Write-Test by:Leftyer dt:2025-11-01 */
func TestJsonWrite(t *testing.T) {
	f := filepath.Join(tmp, "write.json")
	defer os.Remove(f)

	type U struct {
		Name string
		Age  int
	}
	if err := JsonWrite(f, U{Name: "Go", Age: 20}); err != nil {
		t.Fatalf("JsonWrite fail: %v", err)
	}
}

/* Benchmark-Json-Write by:Leftyer dt:2025-11-01 */
func BenchmarkJsonWrite(b *testing.B) {
	f := filepath.Join(tmp, "benchWrite.json")
	defer os.Remove(f)
	type U struct {
		Name string
		Age  int
	}
	v := U{Name: "bench", Age: 99}
	for b.Loop() {
		_ = JsonWrite(f, v)
	}
}
