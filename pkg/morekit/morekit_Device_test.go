/*
 * morekit_Device_test.go Test-difine
 * by:Leftyer
 * dt:2025-11-01
 */
package morekit

import (
	"testing"
)

/* TestUnit-DeviceGetInfo by:Leftyer dt:2025-11-01 */
func TestDeviceGetInfo(t *testing.T) {
	info := DeviceGetInfo()
	if info == nil || info.HostID == "" {
		t.Fatal("DeviceGetInfo returned nil or empty HostID")
	}
}

/* TestBenchmark-DeviceGetInfo by:Leftyer dt:2025-11-01 */
func BenchmarkDeviceGetInfo(b *testing.B) {
	for b.Loop() {
		_ = DeviceGetInfo()
	}
}

/* TestUnit-DeviceGetId by:Leftyer dt:2025-11-01 */
func TestDeviceGetId(t *testing.T) {
	id := DeviceGetId()
	if id == "" {
		t.Fatal("DeviceGetId returned nil or empty id")
	}
}

/* TestBenchmark-DeviceGetId by:Leftyer dt:2025-11-01 */
func BenchmarkDeviceGetId(b *testing.B) {
	for b.Loop() {
		_ = DeviceGetId()
	}
}

/* TestUnit-DeviceGetName by:Leftyer dt:2025-11-01 */
func TestDeviceGetName(t *testing.T) {
	id := DeviceGetName()
	if id == "" {
		t.Fatal("DeviceGetName returned nil or empty name")
	}
}

/* TestBenchmark-DeviceGetName by:Leftyer dt:2025-11-01 */
func BenchmarkDeviceGetName(b *testing.B) {
	for b.Loop() {
		_ = DeviceGetName()
	}
}
