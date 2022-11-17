package main

import (
	"fmt"
	"time"

	"github.com/go-ping/ping"
)

func myPing(host string) {
	p, err := ping.NewPinger(host)
	if err != nil {
		panic(err)
	}
	p.Count = 1
	p.SetPrivileged(true)
	if err = p.Run(); err != nil {
		panic(err)
	}
	stats := p.Statistics() // get send/receive/duplicate/rtt stats
	fmt.Println(host, "OK, latency is", stats.AvgRtt)
}

func main() {
	targets := []string{"yahoo.com", "google.com", "cisco.com", "cern.ch"}

	for _, target := range targets {
		go myPing(target)
	}
	time.Sleep(time.Second * 3) //Wait 3 seconds

}
