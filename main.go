package main

import (
	"exp/ssh"
	"flag"
	"log"
	"os"
)

var (
	USER = flag.String("u", "", "Username")
	PASS = flag.String("p", "", "Password")
	HOST = flag.String("h", "", "Host")
)

func init() {
	flag.Parse()
	if len(*USER) == 0 || len(*HOST) == 0 {
		flag.Usage()
		os.Exit(1)
	}
}

type password string

func (p password) Password(user string) (string, error) {
	return string(p), nil
}

func main() {
	config := &ssh.ClientConfig{
		User: *USER,
		Auth: []ssh.ClientAuth {
			ssh.ClientAuthPassword(password(*PASS)),
		},
	}
	client, err := ssh.Dial("tcp", *HOST, config)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	log.Printf("Connected to %s", client.RemoteAddr())
	cmd, err := client.NewSession()
	if err != nil {
		log.Fatalf("coult not requets new session: %v", err)	
	}
	defer cmd.Close()
	cmd.Stdout = os.Stdout
	if err := cmd.Run("uname -a"); err != nil {
		log.Fatalf("Exec: %v", err)
	}
}
