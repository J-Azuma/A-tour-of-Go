package main

import "fmt"

type IPAddr [4]byte

func (addr IPAddr) String() string {
	var str string
	for i := 0; i < len(addr); i++ {
		str += (fmt.Sprint(addr[i]) + ".")
	}
	return string(str[:len(str)-1])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
