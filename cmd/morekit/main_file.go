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
		cmd  = fs.String("c", "", "sub-command: copy|exist|read|write")
		src  = fs.String("s", "", "source file")
		dst  = fs.String("d", "", "destination file (copy/write only)")
		text = fs.String("t", "", "text content (write only)")
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
		fmt.Println("file exist:", morekit.FileIsExist(*src))

	case "read":
		if *src == "" {
			fmt.Println("file read: need -s path")
			os.Exit(1)
		}
		s, err := morekit.FileRead(*src)
		if err != nil {
			fmt.Println("file read err:", err)
			os.Exit(1)
		}
		fmt.Print(s)

	case "write":
		if *dst == "" || *text == "" {
			fmt.Println("file write: need -d path -t content")
			os.Exit(1)
		}
		if err := morekit.FileWrite(*dst, *text); err != nil {
			fmt.Println("file write err:", err)
			os.Exit(1)
		}
		fmt.Println("file written →", *dst)

	default:
		fmt.Println("file: need -c copy|exist|read|write")
		os.Exit(1)
	}
}

/* Json by:Leftyer dt:2025-11-01 */
func runJson(fs *flag.FlagSet) {
	type sample struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var (
		cmd  = fs.String("c", "", "sub-command: read|write")
		path = fs.String("p", "", "json file path")
	)
	fs.Parse(os.Args[3:])

	switch *cmd {
	case "read":
		if *path == "" {
			fmt.Println("json read: need -p path")
			os.Exit(1)
		}
		v, err := morekit.JsonRead[sample](*path)
		if err != nil {
			fmt.Println("json read err:", err)
			os.Exit(1)
		}
		fmt.Printf("name=%s age=%d\n", v.Name, v.Age)

	case "write":
		if *path == "" {
			fmt.Println("json write: need -p path")
			os.Exit(1)
		}
		if err := morekit.JsonWrite(*path, sample{Name: "Leftyer", Age: 18}); err != nil {
			fmt.Println("json write err:", err)
			os.Exit(1)
		}
		fmt.Println("json written →", *path)

	default:
		fmt.Println("json: need -c read|write")
		os.Exit(1)
	}
}
