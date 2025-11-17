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
	t.Log(`
	========== DeviceGetInfo Begin ==========
	`)
	info := DeviceGetInfo()
	if info == nil || info.HostID == "" {
		t.Fatal(`
		DeviceGetInfo returned nil or empty HostID
		`)
	} else {
		t.Logf(`
		DeviceGetInfo:=>%+v
		`, info)
	}
	t.Log(`
	========== DeviceGetInfo End   ==========
	`)
}

/* TestBenchmark-DeviceGetInfo by:Leftyer dt:2025-11-01 */
func BenchmarkDeviceGetInfo(b *testing.B) {
	for b.Loop() {
		_ = DeviceGetInfo()
	}
}

/* TestUnit-DeviceGetId by:Leftyer dt:2025-11-01 */
func TestDeviceGetId(t *testing.T) {
	t.Log(`
	========== DeviceGetId Begin ==========
	`)
	id := DeviceGetId()
	if id == "" {
		t.Fatal(`
		DeviceGetId returned nil or empty
		`)
	} else {
		t.Logf(`
		DeviceGetId:=>%s
		`, id)
	}
	t.Log(`
	========== DeviceGetId End   ==========
	`)
}

/* TestBenchmark-DeviceGetId by:Leftyer dt:2025-11-01 */
func BenchmarkDeviceGetId(b *testing.B) {
	for b.Loop() {
		_ = DeviceGetId()
	}
}

/* TestUnit-DeviceGetName by:Leftyer dt:2025-11-01 */
func TestDeviceGetName(t *testing.T) {
	t.Log(`
	========== DeviceGetName Begin ==========
	`)
	name := DeviceGetName()
	if name == "" {
		t.Fatal(`
		DeviceGetName returned nil or empty
		`)
	} else {
		t.Logf(`
		DeviceGetName:=>%s
		`, name)
	}
	t.Log(`
	========== DeviceGetName End   ==========
	`)
}

/* TestBenchmark-DeviceGetName by:Leftyer dt:2025-11-01 */
func BenchmarkDeviceGetName(b *testing.B) {
	for b.Loop() {
		_ = DeviceGetName()
	}
}

/* TestUnit-DeviceGetBootTime by:Leftyer dt:2025-11-01 */
func TestDeviceGetBootTime(t *testing.T) {
	t.Log(`
	========== DeviceGetBootTime Begin ==========
	`)
	bt := DeviceGetBootTime()
	if bt == "" {
		t.Fatal(`
		DeviceGetBootTime returned empty
		`)
	} else {
		t.Logf(`
		DeviceGetBootTime:=>%s
		`, bt)
	}
	t.Log(`
	========== DeviceGetBootTime End   ==========
	`)
}

/* TestBenchmark-DeviceGetBootTime by:Leftyer dt:2025-11-01 */
func BenchmarkDeviceGetBootTime(b *testing.B) {
	for b.Loop() {
		_ = DeviceGetBootTime()
	}
}

/* TestUnit-DeviceGetBIOSUUID by:Leftyer dt:2025-11-01 */
func TestDeviceGetBIOSUUID(t *testing.T) {
	t.Log(`
	========== DeviceGetBIOSUUID Begin ==========
	`)
	uuid := DeviceGetBIOSUUID()
	if uuid == "" {
		t.Fatal(`
		DeviceGetBIOSUUID returned empty
		`)
	} else {
		t.Logf(`
		DeviceGetBIOSUUID:=>%s
		`, uuid)
	}
	t.Log(`
	========== DeviceGetBIOSUUID End   ==========
	`)
}

/* TestBenchmark-DeviceGetBIOSUUID by:Leftyer dt:2025-11-01 */
func BenchmarkDeviceGetBIOSUUID(b *testing.B) {
	for b.Loop() {
		_ = DeviceGetBIOSUUID()
	}
}

/* TestUnit-DeviceGetCPUName by:Leftyer dt:2025-11-01 */
func TestDeviceGetCPUName(t *testing.T) {
	t.Log(`
	========== DeviceGetCPUName Begin ==========
	`)
	name := DeviceGetCPUName()
	if name == "" {
		t.Fatal(`
		DeviceGetCPUName returned empty
		`)
	} else {
		t.Logf(`
		DeviceGetCPUName:=>%s
		`, name)
	}
	t.Log(`
	========== DeviceGetCPUName End   ==========
	`)
}

/* TestBenchmark-DeviceGetCPUName by:Leftyer dt:2025-11-01 */
func BenchmarkDeviceGetCPUName(b *testing.B) {
	for b.Loop() {
		_ = DeviceGetCPUName()
	}
}

/* TestUnit-DeviceGetCPUCores by:Leftyer dt:2025-11-01 */
func TestDeviceGetCPUCores(t *testing.T) {
	t.Log(`
	========== DeviceGetCPUCores Begin ==========
	`)
	cores := DeviceGetCPUCores()
	if cores == 0 {
		t.Fatal(`
		DeviceGetCPUCores returned 0
		`)
	} else {
		t.Logf(`
		DeviceGetCPUCores:=>%d
		`, cores)
	}
	t.Log(`
	========== DeviceGetCPUCores End   ==========
	`)
}

/* TestBenchmark-DeviceGetCPUCores by:Leftyer dt:2025-11-01 */
func BenchmarkDeviceGetCPUCores(b *testing.B) {
	for b.Loop() {
		_ = DeviceGetCPUCores()
	}
}

/* TestUnit-DeviceGetMemoryGB by:Leftyer dt:2025-11-01 */
func TestDeviceGetMemoryGB(t *testing.T) {
	t.Log(`
	========== DeviceGetMemoryGB Begin ==========
	`)
	gb := DeviceGetMemoryGB()
	if gb == 0 {
		t.Fatal(`
		DeviceGetMemoryGB returned 0
		`)
	} else {
		t.Logf(`
		DeviceGetMemoryGB:=>%.2f GB
		`, gb)
	}
	t.Log(`
	========== DeviceGetMemoryGB End   ==========
	`)
}

/* TestBenchmark-DeviceGetMemoryGB by:Leftyer dt:2025-11-01 */
func BenchmarkDeviceGetMemoryGB(b *testing.B) {
	for b.Loop() {
		_ = DeviceGetMemoryGB()
	}
}

/* TestUnit-DeviceGetDiskTotalGB by:Leftyer dt:2025-11-01 */
func TestDeviceGetDiskTotalGB(t *testing.T) {
	t.Log(`
	========== DeviceGetDiskTotalGB Begin ==========
	`)
	gb := DeviceGetDiskTotalGB()
	if gb == 0 {
		t.Fatal(`
		DeviceGetDiskTotalGB returned 0
		`)
	} else {
		t.Logf(`
		DeviceGetDiskTotalGB:=>%.2f GB
		`, gb)
	}
	t.Log(`
	========== DeviceGetDiskTotalGB End   ==========
	`)
}

/* TestBenchmark-DeviceGetDiskTotalGB by:Leftyer dt:2025-11-01 */
func BenchmarkDeviceGetDiskTotalGB(b *testing.B) {
	for b.Loop() {
		_ = DeviceGetDiskTotalGB()
	}
}

/* TestUnit-DeviceGetIPList by:Leftyer dt:2025-11-01 */
func TestDeviceGetIPList(t *testing.T) {
	t.Log(`
	========== DeviceGetIPList Begin ==========
	`)
	list := DeviceGetIPList()
	if len(list) == 0 {
		t.Fatal(`
		DeviceGetIPList returned empty
		`)
	} else {
		t.Logf(`
		DeviceGetIPList:=>%v
		`, list)
	}
	t.Log(`
	========== DeviceGetIPList End   ==========
	`)
}

/* TestBenchmark-DeviceGetIPList by:Leftyer dt:2025-11-01 */
func BenchmarkDeviceGetIPList(b *testing.B) {
	for b.Loop() {
		_ = DeviceGetIPList()
	}
}

/* TestUnit-DeviceGetMACList by:Leftyer dt:2025-11-01 */
func TestDeviceGetMACList(t *testing.T) {
	t.Log(`
	========== DeviceGetMACList Begin ==========
	`)
	list := DeviceGetMACList()
	if len(list) == 0 {
		t.Fatal(`
		DeviceGetMACList returned empty
		`)
	} else {
		t.Logf(`
		DeviceGetMACList:=>%v
		`, list)
	}
	t.Log(`
	========== DeviceGetMACList End   ==========
	`)
}

/* TestBenchmark-DeviceGetMACList by:Leftyer dt:2025-11-01 */
func BenchmarkDeviceGetMACList(b *testing.B) {
	for b.Loop() {
		_ = DeviceGetMACList()
	}
}

/* TestUnit-DeviceGetDefaultIP by:Leftyer dt:2025-11-01 */
func TestDeviceGetDefaultIP(t *testing.T) {
	t.Log(`
	========== DeviceGetDefaultIP Begin ==========
	`)
	ip := DeviceGetDefaultIP()
	if ip == "" {
		t.Fatal(`
		DeviceGetDefaultIP returned empty
		`)
	} else {
		t.Logf(`
		DeviceGetDefaultIP:=>%s
		`, ip)
	}
	t.Log(`
	========== DeviceGetDefaultIP End   ==========
	`)
}

/* TestBenchmark-DeviceGetDefaultIP by:Leftyer dt:2025-11-01 */
func BenchmarkDeviceGetDefaultIP(b *testing.B) {
	for b.Loop() {
		_ = DeviceGetDefaultIP()
	}
}
