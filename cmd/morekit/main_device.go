/*
 * main_device.go main-difine
 * by:Leftyer
 * dt:2025-11-01
 */
package main

import (
	"flag"
	"fmt"

	"github.com/LeftyerPro/more-kit/pkg/morekit"
)

/* Device by:Leftyer dt:2025-11-01 */
func runDevice(fs *flag.FlagSet) {
	v := func(name string) bool { return fs.Lookup(name).Value.(flag.Getter).Get().(bool) }

	if fs.NFlag() == 0 {
		v = func(name string) bool { return name == "info" }
	}

	if v("id") {
		fmt.Println("Device ID  :", morekit.DeviceGetId())
	}
	if v("name") {
		fmt.Println("Device Name:", morekit.DeviceGetName())
	}
	if v("info") {
		if info := morekit.DeviceGetInfo(); info != nil {
			fmt.Printf("%+v\n", *info)
		}
	}
	if v("boot") {
		fmt.Println("Boot Time  :", morekit.DeviceGetBootTime())
	}
	if v("uuid") {
		fmt.Println("BIOS UUID  :", morekit.DeviceGetBIOSUUID())
	}
	if v("cpu") {
		fmt.Println("CPU Model  :", morekit.DeviceGetCPUName())
	}
	if v("cores") {
		fmt.Println("CPU Cores  :", morekit.DeviceGetCPUCores())
	}
	if v("mem") {
		fmt.Printf("Memory     : %.2f GB\n", morekit.DeviceGetMemoryGB())
	}
	if v("disk") {
		fmt.Printf("Disk Total : %.2f GB\n", morekit.DeviceGetDiskTotalGB())
	}
	if v("ips") {
		fmt.Println("IP List    :", morekit.DeviceGetIPList())
	}
	if v("macs") {
		fmt.Println("MAC List   :", morekit.DeviceGetMACList())
	}
	if v("ip") {
		fmt.Println("Default IP :", morekit.DeviceGetDefaultIP())
	}
}

/* Device-flags by:Leftyer dt:2025-11-01 */
func defineDeviceFlags(fs *flag.FlagSet) {
	fs.Bool("id", false, "get device HostID")
	fs.Bool("name", false, "get hostname")
	fs.Bool("info", false, "get full host.InfoStat")
	fs.Bool("boot", false, "get boot time")
	fs.Bool("uuid", false, "get BIOS UUID")
	fs.Bool("cpu", false, "get CPU model")
	fs.Bool("cores", false, "get CPU cores")
	fs.Bool("mem", false, "get memory GB")
	fs.Bool("disk", false, "get disk total GB")
	fs.Bool("ips", false, "get IPv4 list")
	fs.Bool("macs", false, "get MAC list")
	fs.Bool("ip", false, "get default outbound IP")
}
