package main

import (
	"fmt"
	"log"

	gomiko "github.com/Ali-aqrabawi/gomiko/pkg"
)

func main() {
	device := gomiko.NewDevice(
		"10.0.4.1", "router1", "router1", "juniper",
	)
	if err := device.Connect(); err != nil {
		log.Fatal(err)
	}

	result, _ := device.SendCommand("show version")
	device.Disconnect()
	fmt.Println(result)

}
