package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/yahoo/vssh"
)

func main() {
	vs := vssh.New().Start()
	config := vssh.GetConfigUserPass("netlab", "netlab")
	vs.AddClient(
		"10.0.4.1:22", config, vssh.SetMaxSessions(1),
	)
	vs.Wait()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	timeout, _ := time.ParseDuration("10s")
	rChannel := vs.Run(ctx, "uptime", timeout)

	for resp := range rChannel {
		if err := resp.Err(); err != nil {
			log.Println(err)
			continue
		}
		outTxt, _, _ := resp.GetText(vs)
		fmt.Println(outTxt)
	}
}
