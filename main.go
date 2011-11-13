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
	kc := new(keychain)
	if err := kc.LoadPEM(os.Getenv("HOME") + "/.ssh/id_rsa"); err != nil {
		log.Fatal(err)
	}
	config := &ssh.ClientConfig{
		User: *USER,
		Auth: []ssh.ClientAuth{
			ssh.ClientAuthPublickey(kc),
			ssh.ClientAuthPassword(password(*PASS)),
		},
	}
	client, err := ssh.Dial("tcp", *HOST, config)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	log.Printf("Connected to %s", client.RemoteAddr())
	// open a few channels to bring the id and peerid's out of sync
	if _, err := client.NewSession(); err != nil {
		log.Fatal(err)
	}
	if _, err := client.NewSession(); err != nil {
		log.Fatal(err)
	}
	if _, err := client.NewSession(); err != nil {
		log.Fatal(err)
	}

	shell, err := client.NewSession()
	shell.Stdin = os.Stdin
	shell.Stdout = os.Stdout
	shell.Stderr = os.Stderr
	if err != nil {
		log.Fatal(err)
	}
	if err := shell.RequestPty("vt100", 80, 25); err != nil {
		log.Fatal(err)
	}
	if err := shell.Shell(); err != nil {
		log.Fatal(err)
	}
	log.Println("Shell opened")
	select {}
}
