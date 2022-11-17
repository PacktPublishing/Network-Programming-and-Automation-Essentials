package main

import (
	"fmt"
	"sync"

	"github.com/go-ping/ping"
)

func myping(host string, wg *sync.WaitGroup) {
	defer wg.Done()
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
	var targets = []string{"yahoo.com", "google.com", "cisco.com", "cern.ch"}

	var wg sync.WaitGroup
	wg.Add(len(targets))
	for _, target := range targets {
		go myping(target, &wg)
	}
	wg.Wait()

}
