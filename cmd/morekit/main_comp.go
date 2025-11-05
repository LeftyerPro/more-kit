/*
 * main_comp.go main-difine
 * by:Leftyer
 * dt:2025-11-01
 */
package main

import (
	"flag"
	"log"
	"os"

	"github.com/LeftyerPro/more-kit/pkg/morekit"
)

/* Comp by:Leftyer dt:2025-11-01 */
func runComp(fs *flag.FlagSet) {
	var (
		in  = fs.String("i", "logo.png", "input file")
		out = fs.String("o", "logo.webp", "output file")
	)
	fs.Parse(os.Args[3:])
	if err := morekit.CompImage(*in, *out, 3); err != nil {
		log.Fatal(err)
	}
	println("compressed →", *out)
}
