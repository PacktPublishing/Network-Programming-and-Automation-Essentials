package main

import (
	"fmt"
	"log"

	"github.com/scrapli/scrapligo/driver/generic"
	"github.com/scrapli/scrapligo/driver/options"
)

func main() {
	target, err := generic.NewDriver(
		"10.0.4.1",
		options.WithAuthNoStrictKey(),
		options.WithAuthUsername("netlab"),
		options.WithAuthPassword("netlab"),
	)
	if err != nil {
		log.Fatalf("Failed to create target: %+v\n", err)
	}

	if err = target.Open(); err != nil {
		log.Fatalf("Failed to open: %+v\n", err)
	}

	output, err := target.Channel.SendInput("uptime")
	if err != nil {
		log.Fatalf("Failed to send command: %+v\n", err)
	}

	fmt.Println(string(output))

}
