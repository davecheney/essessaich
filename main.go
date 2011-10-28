package main

import (
	"exp/ssh"
	"flag"
	"io"
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
	if len(*USER) == 0 || len(*PASS) == 0 || len(*HOST) == 0 {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	config := &ssh.ClientConfig{
		User:                  *USER,
		Password:              *PASS,
	}
	client, err := ssh.Dial("tcp", *HOST, config)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	log.Printf("Connected to %s", client.RemoteAddr())
	// open a few channels to bring the id and peerid's out of sync
	if _, err := client.NewSession() ; err != nil {
		log.Fatal(err)
	}
        if _, err := client.NewSession() ; err != nil {
                log.Fatal(err)
        }
        if _, err := client.NewSession() ; err != nil {
                log.Fatal(err)
        }

        shell, err := client.NewSession()
        if err != nil {
                log.Fatal(err)
        }
	if err := shell.RequestPty("vt100", 80, 25); err != nil {
		log.Fatal(err)
	}
        if err := shell.Shell() ; err != nil {
                log.Fatal(err)
        }
	log.Println("Shell opened")
	go io.Copy(os.Stderr, shell.Stderr)
	go io.Copy(os.Stdin, shell.Stdout)
	if _, err := io.Copy(shell.Stdin, os.Stdout); err != nil {
		log.Fatal(err)
	}

}
