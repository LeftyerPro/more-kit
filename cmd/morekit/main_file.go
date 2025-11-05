/*
 * main_file.go main-difine
 * by:Leftyer
 * dt:2025-11-01
 */
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/LeftyerPro/more-kit/pkg/morekit"
)

/* Folder by:Leftyer dt:2025-11-01 */
func runFolder(fs *flag.FlagSet) {
	var (
		cmd = fs.String("c", "", "sub-command: copy|exist")
		src = fs.String("s", "", "source folder")
		dst = fs.String("d", "", "destination folder (copy only)")
		clr = fs.Int("clear", 0, "clear if exist (exist only)")
	)
	fs.Parse(os.Args[3:])

	switch *cmd {
	case "copy":
		if *src == "" || *dst == "" {
			fmt.Println("folder copy: need -s src -d dst")
			os.Exit(1)
		}
		if err := morekit.FolderCpoy(*src, *dst); err != nil {
			fmt.Println("folder copy err:", err)
			os.Exit(1)
		}
		fmt.Println("folder copied →", *dst)

	case "exist":
		if *src == "" {
			fmt.Println("folder exist: need -s path")
			os.Exit(1)
		}
		yes := morekit.FolderIsExist(*src, *clr)
		fmt.Println("folder exist:", yes)

	default:
		fmt.Println("folder: need -c copy|exist")
		os.Exit(1)
	}
}

/* File by:Leftyer dt:2025-11-01 */
func runFile(fs *flag.FlagSet) {
	var (
		cmd = fs.String("c", "", "sub-command: copy|exist")
		src = fs.String("s", "", "source file")
		dst = fs.String("d", "", "destination file (copy only)")
	)
	fs.Parse(os.Args[3:])

	switch *cmd {
	case "copy":
		if *src == "" || *dst == "" {
			fmt.Println("file copy: need -s src -d dst")
			os.Exit(1)
		}
		if err := morekit.FileCpoy(*src, *dst); err != nil {
			fmt.Println("file copy err:", err)
			os.Exit(1)
		}
		fmt.Println("file copied →", *dst)

	case "exist":
		if *src == "" {
			fmt.Println("file exist: need -s path")
			os.Exit(1)
		}
		yes := morekit.FileIsExist(*src)
		fmt.Println("file exist:", yes)

	default:
		fmt.Println("file: need -c copy|exist")
		os.Exit(1)
	}
}
