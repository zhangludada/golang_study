package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("0.0.0.0"), Port: 30000})
	if err != nil {
		fmt.Println(err)
	}
	defer listen.Close()

	for {
		var buf [1024]byte
		n, addr, err := listen.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("recStr:", string(buf[:n]))
		_, err = listen.WriteToUDP(buf[:], addr)
		if err != nil {
			fmt.Println(err)
		}
	}
}
