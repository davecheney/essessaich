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
}

func main() {
	config := &ssh.ClientConfig{
		User: *USER,
		Password: *PASS,
		SupportedKexAlgos:[]string{"diffie-hellman-group14-sha1"}, 
		SupportedHostKeyAlgos:[]string{"ssh-rsa"}, 
		SupportedCiphers:[]string{"aes128-ctr"}, 
		SupportedMACs:[]string{"hmac-sha1-96"}, 
		SupportedCompressions:[]string{"none"},
	}
	client, err := ssh.Dial(*HOST, config)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	shell, err := client.OpenChan("session")
	if err != nil {
		log.Fatal(err)
	}
	if err := shell.Shell(); err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			buf := make([]byte, 4096)
			read, err := os.Stdin.Read(buf)
			if err != nil {
				log.Fatal(err)
			}
			shell.Write(buf[:read])
		}
	}()
	for {
		buf := make([]byte, 4096)
		read, err := shell.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		os.Stdout.Write(buf[:read])
	}
	
}
