/*
 * main_version.go main-difine
 * by:Leftyer
 * dt:2025-11-01
 */
package main

import "flag"

const version = "1.0.0"

/* Version by:Leftyer dt:2025-11-01 */
func runVersion(fs *flag.FlagSet) {
	fs.Parse([]string{})
	println("more-kit version", version)
}

/* Help by:Leftyer dt:2025-11-01 */
func runHelp(fs *flag.FlagSet) {
	fs.Parse([]string{})
	println("usage: morekit <command> [<args>]")
	println("commands:")
	println("  version                    show version")
	println("  help                       show this help")
	println("  device                     show device id & name")
	println("  comp -i <in> -o <out>   lossless comp to webp")
	println("  folder -c copy|exist -s <src> [-d <dst>] [-clear 0|1]   copy or check folder")
	println("  file   -c copy|exist -s <src> [-d <dst>]               copy or check file")
}
