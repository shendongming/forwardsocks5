package main

import (
	"flag"
	"fmt"
	"os"
	"socks5server"
)

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}
func main() {
	name := flag.String("name", "", "Input server name")
	port := flag.Int("socks5", 1086, "Listen socks5 proxy server port")
	port2 := flag.Int("work", 1088, "Listen work server port")
	flag.Parse()
	if len(*name) == 0 {
		Usage()
		return
	}
	addr := fmt.Sprintf(":%d" ,*port)
	addr2 := fmt.Sprintf(":%d" ,*port2)
	ser := socks5server.NewServer(addr,addr2)
	println(*name)
	println("socks5 addr",addr,"workd addr:",addr2)
	err:=ser.ListenAndServe()
	println("error",err)
}