package main

import (
	"flag"
	"fmt"
	"os"
	"work"
)

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}
func main() {
	addr := flag.String("addr", "", "Input server address example: 192.168.2.143:1088 ")

	flag.Parse()
	if len(*addr) == 0 {
		Usage()
		return
	}

	work.Connect(*addr)
}