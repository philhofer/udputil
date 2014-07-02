package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

var (
	bindAddr string
)

func init() {
	flag.StringVar(&bindAddr, "bind", ":65000", "UDP local listen address")
}

func main() {
	flag.Parse()

	addr, err := net.ResolveUDPAddr("udp", bindAddr)
	if err != nil {
		fmt.Printf("Couldn't resolve bind address: %s\n", err.Error())
		os.Exit(1)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Printf("Listen error: %s", err.Error())
		os.Exit(1)
	}
	log.Printf("Listening on %s...", conn.LocalAddr().String())
	buf := [1024]byte{}
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt, os.Kill)
	for {
		select {
		case <-sigchan:
			goto exit
		default:
			// Timeout forces us to poll on sigchan regularly; otherwise receive blocks
			conn.SetReadDeadline(time.Now().Add(100 * time.Millisecond))
			n, addr, err := conn.ReadFromUDP(buf[0:])
			if err != nil {
				if neterr, ok := err.(net.Error); ok {
					if neterr.Timeout() {
						continue
					}
				}
				log.Fatal(err)
			}
			log.Printf("From %s: %q", addr.String(), buf[0:n])
		}
	}
exit:
	conn.Close()
	log.Println("Exited normally.")
}
