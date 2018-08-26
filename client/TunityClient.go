package main

import (
	"net"
	"fmt"
)

func main() {
	ua, err := net.ResolveUDPAddr("udp","localhost:4567")
	if err != nil {
		fmt.Println(err)
		return
	}

	uc, err := net.DialUDP("udp",nil, ua)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer uc.Close()
	buffer := make([]byte, 1024)
	buffer[0] = 255
	uc.Write(buffer)


}
