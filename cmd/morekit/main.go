/*
 * main.go main-difine
 * by:Leftyer
 * dt:2025-11-01
 */
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/LeftyerPro/more-kit/pkg/morekit"
)

var ver = flag.Bool("v", false, "version")

const version = "1.0.0"

func main() {
	flag.Parse()
	if *ver {
		fmt.Println(version)
		return
	}
	if err := morekit.CompImage(`F:\Test\input\logo.png`, `F:\Test\output\logo.png`, 3); err != nil {
		log.Fatal(err)
	}
}
