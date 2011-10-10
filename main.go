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
		SupportedKexAlgos:     []string{"diffie-hellman-group14-sha1"},
		SupportedHostKeyAlgos: []string{"ssh-rsa"},
		SupportedCiphers:      []string{"aes128-ctr"},
		SupportedMACs:         []string{"hmac-sha1-96"},
		SupportedCompressions: []string{"none"},
	}
	client, err := ssh.Dial(*HOST, config)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	log.Printf("Connected to %s", client.RemoteAddr())
	// open a few channels to bring the id and peerid's out of sync
	client.OpenChan("session") ; client.OpenChan("session")
	shell, err := client.OpenChan("session")
	if err != nil {
		log.Fatal(err)
	}
	if err := shell.Ptyreq("vt100", 80, 25); err != nil {
		log.Fatal(err)
	}
	stdin, stdout, stderr, err := shell.Shell()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Shell opened")
	go io.Copy(os.Stderr, stderr)
	go io.Copy(os.Stdin, stdout)
	if _, err := io.Copy(stdin, os.Stdout); err != nil {
		log.Fatal(err)
	}

}
