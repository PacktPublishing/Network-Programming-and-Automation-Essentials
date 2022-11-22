package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/go-ping/ping"
)

func probe(host string, w http.ResponseWriter, wg *sync.WaitGroup) {
	defer wg.Done()
	p, err := ping.NewPinger(host)
	if err != nil {
		fmt.Fprintf(w, "error ping creation: %v\n", err)
		return
	}
	p.Count = 1
	p.Timeout = time.Second * 2
	p.SetPrivileged(true)
	if err = p.Run(); err != nil {
		fmt.Fprintf(w, "error ping sent: %v\n", err)
		return
	}
	stats := p.Statistics()
	if stats.PacketLoss == 0 {
		fmt.Fprintf(w, "%s latency is %s\n", host, stats.AvgRtt)
	} else {
		fmt.Fprintf(w, "%s no response timeout\n", host)
	}
}

func probeTargets(w http.ResponseWriter, r *http.Request) {
	httpTargets := r.URL.Query().Get("targets")
	targets := strings.Split(httpTargets, ",")
	if len(httpTargets) == 0 {
		fmt.Fprintf(w, "No targets to send ICMP\n")
		return
	}

	var wg sync.WaitGroup
	wg.Add(len(targets))
	for _, target := range targets {
		log.Println("requested ICMP probe for", target)
		go probe(target, w, &wg)
	}
	wg.Wait()
}

func main() {
	http.HandleFunc("/latency", probeTargets)
	log.Fatal(http.ListenAndServe(":9900", nil))
}
