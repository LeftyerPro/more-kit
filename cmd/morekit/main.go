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

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: morekit <command> [<args>]")
		fmt.Println("commands: version ,device ,comp")
		os.Exit(1)
	}
	cmdName := os.Args[1]
	cmd, ok := cmds[cmdName]
	if !ok {
		fmt.Printf("unknown command: %s\n", cmdName)
		os.Exit(1)
	}
	cmd.Parse(os.Args[2:])
	switch cmdName {
	case "", "version":
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
