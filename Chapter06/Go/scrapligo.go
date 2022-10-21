package main

import (
	"fmt"
	"log"

	"github.com/scrapli/scrapligo/driver/generic"
	"github.com/scrapli/scrapligo/driver/options"
	"github.com/scrapli/scrapligo/logging"
)

func myLoggerFunc(x ...interface{}) {
	// not sure why you would want to do this, but this is just an example logger function you
	// could use :)
	if len(x) > 0 {
		// pro tip... its almost certainly going to be a string, but this is just to show how
		// you can add logger functions so who cares!
		fmt.Printf("got a log item of type %T\n", x[0])
	}
}

func main() {
	li, err := logging.NewInstance(
		logging.WithLevel(logging.Debug),
		logging.WithLogger(log.Print),
		logging.WithLogger(myLoggerFunc),
	)
	if err != nil {
		fmt.Printf("failed to logging instance; error: %+v\n", err)

		return
	}

	d, err := generic.NewDriver(
		"sandbox-iosxe-latest-1.cisco.com",
		options.WithAuthNoStrictKey(),
		options.WithAuthUsername("developer"),
		options.WithAuthPassword("C1sco12345"),
		/*
			"10.0.4.1",
			options.WithAuthNoStrictKey(),
			options.WithAuthUsername("netlab"),
			options.WithAuthPassword("netlab"),
		*/
		options.WithLogger(li),
	)
	fmt.Printf("Got:%+v\n", d)
	if err != nil {
		fmt.Printf("failed to create driver; error: %+v\n", err)
		return
	}

	err = d.Open()
	if err != nil {
		fmt.Printf("failed to open driver; error: %+v\n", err)
		return
	}
	defer d.Close()

	fmt.Println("Feching the prompt")

	// fetch the prompt
	prompt, err := d.GetPrompt()
	if err != nil {
		fmt.Printf("failed to get prompt; error: %+v\n", err)
		return
	}

	fmt.Printf("found prompt: %s\n\n\n", prompt)

	// send some input
	output, err := d.Channel.SendInput("uptime")
	if err != nil {
		fmt.Printf("failed to send input to device; error: %+v\n", err)
		return
	}

	fmt.Printf("output received (SendInput):\n %s\n\n\n", output)

}
