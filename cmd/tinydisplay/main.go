package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/sago35/tinydisplay/server"
)

func main() {
	addr := flag.String("address", "127.0.0.1", "listen address")
	port := flag.Int("port", 9812, "listen port")
	size := flag.String("size", "320x240", "display size (ex: 320x240)")
	flag.Parse()

	w, h := 0, 0
	n, err := fmt.Sscanf(*size, "%dx%d", &w, &h)
	if err != nil {
		log.Fatal(err)
	}
	if n != 2 {
		log.Fatal("size format error : n != 2 (%d)", n)
	}

	fmt.Printf("tcp:%s:%d\n", *addr, *port)
	fmt.Printf("%dx%d\n", w, h)
	err = run(*addr, *port, w, h)
	if err != nil {
		log.Fatal(err)
	}
}

func run(addr string, port, w, h int) error {
	server := server.NewServer(w, h)
	rpc.Register(server)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		return err
	}
	go http.Serve(l, nil)

	server.ShowAndRun(nil, nil)
	return nil
}
