package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	c, err := net.DialUDP("udp4", nil, &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 30000})
	if err != nil {
		fmt.Println(err)
	}
	defer c.Close()

	input := bufio.NewReader(os.Stdin)
	for {
		s, err := input.ReadString('\n')
		if err != nil {
			fmt.Println(err)

		}
		_, err = c.Write([]byte(s))
		if err != nil {
			fmt.Println(err)
		}

		var buf [1024]byte
		n, _, err := c.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(buf[:n]))
	}
}
