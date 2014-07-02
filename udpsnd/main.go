package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
  "io"
)

var (
	saddr string
)

func init() {
	flag.StringVar(&saddr, "bind", ":65000", "Address to publish to")
}

func main() {
	flag.Parse()

	remote, err := net.ResolveUDPAddr("udp", saddr)
	if err != nil {
		fmt.Printf("Couldn't resolve remote address: %s\n", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, remote)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	conn.SetWriteBuffer(2048)

	fmt.Printf("Sending from %s, sending to %s\n...", conn.LocalAddr().String(), remote.String())
	bin := bufio.NewReaderSize(os.Stdin, 1024)
	for {
		line, _, err := bin.ReadLine()
		if err != nil {
      if err == io.EOF {
        os.Exit(0)
      }
			fmt.Println(err)
			os.Exit(1)
		}
		_, err = conn.Write(line)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	}
}
