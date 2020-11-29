package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func ip2int(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

func int2ip(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}

func main() {
	ip := net.ParseIP("192.168.8.44")
	num := ip2int(ip)
	fmt.Printf("ip2int %s => %d\n", "192.168.8.44", num)
	fmt.Printf("int2ip %d => %s", num, int2ip(num))
}