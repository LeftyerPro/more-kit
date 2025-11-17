/*
 * main_version.go main-difine
 * by:Leftyer
 * dt:2025-11-01
 */
package main

import (
	"flag"
	"fmt"
)

const version = "1.0.0"

/* Version by:Leftyer dt:2025-11-01 */
func runVersion(fs *flag.FlagSet) {
	if fs != nil {
		fs.Parse([]string{})
	}
	fmt.Println("more-kit version", version)
}

/* Help by:Leftyer dt:2025-11-01 */
func runHelp(fs *flag.FlagSet) {
	if fs != nil {
		fs.Parse([]string{})
	}

	fmt.Println("usage: morekit <command> [<args>]")
	fmt.Println()
	fmt.Println("global flags:")
	fmt.Println("  -h, help      show this help")
	fmt.Println("  -v, version   show version")
	fmt.Println()
	fmt.Println("commands:")
	fmt.Println("  device   device utilities")
	fmt.Println("  comp     lossless image compression")
	fmt.Println("  folder   folder copy / check")
	fmt.Println("  file     file read / write / copy / check")
	fmt.Println("  json     json read / write")
	fmt.Println()
	fmt.Println("device flags:")
	fmt.Println("  -id        device HostID")
	fmt.Println("  -name      hostname")
	fmt.Println("  -info      full host.InfoStat")
	fmt.Println("  -boot      boot time")
	fmt.Println("  -uuid      BIOS UUID")
	fmt.Println("  -cpu       CPU model")
	fmt.Println("  -cores     CPU cores")
	fmt.Println("  -mem       memory GB")
	fmt.Println("  -disk      disk total GB")
	fmt.Println("  -iplist    IPv4 list")
	fmt.Println("  -maclist   MAC list")
	fmt.Println("  -ip        default outbound IP")
	fmt.Println()
	fmt.Println("examples:")
	fmt.Println("  morekit -h")
	fmt.Println("  morekit device -info")
	fmt.Println("  morekit comp -i input.png -o output.webp")
	fmt.Println("  morekit folder -c copy -s src -d dst")
	fmt.Println("  morekit file  -c read  -p path")
	fmt.Println("  morekit json  -c read  -p path")
}
