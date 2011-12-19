package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io"
	"io/ioutil"
)

type keychain struct {
	keys []*rsa.PrivateKey
}

func (k *keychain) Key(i int) (interface{}, error) {
	if i < 0 || i >= len(k.keys) {
		return nil, nil
	}
	return k.keys[i].PublicKey, nil
}

// Sign returns a signature of the given data using the i'th key.
func (k *keychain) Sign(i int, rand io.Reader, data []byte) (sig []byte, err error) {
	pk := k.keys[i]
	h := sha1.New()
	h.Write(data)
	hh := h.Sum(nil)
	return rsa.SignPKCS1v15(rand, pk, crypto.SHA1, hh)
}

func (k *keychain) LoadPEM(file string) error {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	block, _ := pem.Decode(buf)
	if block == nil {
		return errors.New("ssh: no key found")
	}
	r, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return err
	}
	k.keys = append(k.keys, r)
	return nil
}
