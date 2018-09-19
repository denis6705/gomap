package main

import (
	"fmt"
	"net"
	"time"

	ping "github.com/digineo/go-ping"
)

var (
	timeout   = time.Second
	ping_time time.Duration
	attempts  uint = 3

	remoteAddr *net.IPAddr
	pinger     *ping.Pinger
)

func main() {
	pinger, _ = ping.New("0.0.0.0", "")
	remoteAddr, _ = net.ResolveIPAddr("ip4", "google.com")
	ping_time, _ = pinger.PingAttempts(remoteAddr, timeout, 3)
	fmt.Println("time: ", ping_time)

}
