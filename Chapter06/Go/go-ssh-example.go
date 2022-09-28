package main

import (
	"bytes"
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

func main() {
	host := "10.0.4.1"

	config := &ssh.ClientConfig{
		User:            "netlab",
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password("netlab"),
		},
	}
	conn, err := ssh.Dial("tcp", host+":22", config)
	if err != nil {
		log.Fatalf("Dial failed: %s", err)
	}
	session, err := conn.NewSession()
	if err != nil {
		log.Fatalf("NewSession failed: %s", err)
	}

	var buff bytes.Buffer
	session.Stdout = &buff
	if err := session.Run("uptime"); err != nil {
		log.Fatalf("Run failed: %s", err)
	}

	fmt.Println(buff.String())
}
