/*
 * main.go main-difine
 * by:Leftyer
 * dt:2025-11-01
 */
package main

import (
	"flag"
	"fmt"
	"os"
)

var cmds = map[string]*flag.FlagSet{
	"version": flag.NewFlagSet("version", flag.ExitOnError),
	"help":    flag.NewFlagSet("help", flag.ExitOnError),
	"device":  flag.NewFlagSet("device", flag.ExitOnError),
	"comp":    flag.NewFlagSet("comp", flag.ExitOnError),
	"folder":  flag.NewFlagSet("folder", flag.ExitOnError),
	"file":    flag.NewFlagSet("file", flag.ExitOnError),
	"json":    flag.NewFlagSet("json", flag.ExitOnError),
}

func init() {
	defineDeviceFlags(cmds["device"])
}

func main() {
	if len(os.Args) < 2 {
		runHelp(nil)
		return
	}
	first := os.Args[1]
	switch first {
	case "-h", "help":
		runHelp(nil)
		return
	case "-v", "version":
		runVersion(nil)
		return
	}

	cmd, ok := cmds[first]
	if !ok {
		fmt.Printf("unknown command: %s\n", first)
		os.Exit(1)
	}
	cmd.Parse(os.Args[2:])
	switch first {
	case "version":
		runVersion(cmd)
	case "help":
		runHelp(cmd)
	case "device":
		runDevice(cmd)
	case "comp":
		runComp(cmd)
	case "folder":
		runFolder(cmd)
	case "file":
		runFile(cmd)
	case "json":
		runJson(cmd)
	}
}
