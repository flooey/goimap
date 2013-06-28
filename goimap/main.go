package main

import (
	"flag"
	"fmt"
	"net"
	"github.com/flooey/goimap/parser"
)

type callback struct { }

func (cb *callback) Noop(tag []byte) {
	fmt.Print("Noop, tag: ")
	fmt.Print(tag)
	fmt.Print("\n")
}

func (cb *callback) Bad(tag []byte) {
	fmt.Print("Bad command, tag: ")
	fmt.Print(tag)
	fmt.Print("\n")
}

func handleConnection(conn net.Conn) {
	fmt.Printf("Connection established!\n")
	buf := make([]byte, 10000)
	cb := new(callback)
	p := parser.MakeParser(cb)
	for {
		l, err := conn.Read(buf)
		if l == 0 || err != nil {
			conn.Close()
			fmt.Printf("Connection closed.\n")
			return
		}
		p.Parse(buf[0:l])
	}
}

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		fmt.Printf("No port\n")
		return
	}
	ln, err := net.Listen("tcp", ":" + flag.Args()[0])
	if err != nil {
		return
	}
	fmt.Printf("Listening.\n")
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn)
	}
}