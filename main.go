package main

import (
	"flag"
	"fmt"
	"net"
	"time"

	ping "github.com/digineo/go-ping"
)

var (
	timeout    = time.Second
	ping_time  time.Duration
	attempts   uint = 3
	host       string
	remoteAddr *net.IPAddr
	pinger     *ping.Pinger
)

func main() {
	flag.StringVar(&host, "host", host, "usage")
	flag.Parse()
	pinger, _ = ping.New("0.0.0.0", "")
	remoteAddr, _ = net.ResolveIPAddr("ip4", host)
	if remoteAddr != nil {
		ping_time, _ = pinger.PingAttempts(remoteAddr, timeout, 3)
		fmt.Println("host: ", host, " time: ", ping_time)
	} else {
		fmt.Println("ipaddress is bad!")
	}

}
