/*
 * main_device.go main-difine
 * by:Leftyer
 * dt:2025-11-01
 */
package main

import (
	"flag"
	"os"

	"github.com/LeftyerPro/more-kit/pkg/morekit"
)

/* Device by:Leftyer dt:2025-11-01 */
func runDevice(fs *flag.FlagSet) {
	fs.Parse(os.Args[3:])
	println("Device ID  :", morekit.DeviceGetId())
	println("Device Name:", morekit.DeviceGetName())
}
