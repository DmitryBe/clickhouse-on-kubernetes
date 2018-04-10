package main

import (
	"log"
	"net"
)

// GetLocalIP returns local node/pod ip
func GetLocalIP() net.IP {

	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
		panic("cannot get local ip")
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}
