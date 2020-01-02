package main

import (
	"log"
	"net"
	"time"
)

const (
	serverAddr      = "224.0.0.1:9999"
	maxDatagramSize = 8192
)

func ping(addrstr string) {

	addr, err := net.ResolveUDPAddr("udp", addrstr)
	if err != nil {
		log.Fatal(err)
	}

	sock, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		_, _ = sock.Write([]byte("hello world\n"))
		time.Sleep(time.Second)
	}
}

func server(addrstr string) {
	addr, err := net.ResolveUDPAddr("udp", addrstr)
	if err != nil {
		log.Fatal(err)
	}
	sock, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}
	_ = sock.SetReadBuffer(maxDatagramSize)
	for {
		buffer := make([]byte, maxDatagramSize)
		n, src, err := sock.ReadFromUDP(buffer)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(n, "bytes read from ", src)
		log.Printf(string(buffer[:n]))
	}
}

func main() {

	go ping(serverAddr)
	server(serverAddr)

}
