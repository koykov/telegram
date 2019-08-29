package api

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"os"
)

// Read and parse RSA public key by given path.
func ParseRSAPublicKey(filePath string) (*rsa.PublicKey, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, errors.New("file not exists: " + filePath)
	}
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(bytes)
	if block == nil {
		return nil, errors.New("wrong RSA public key format, pem expected")
	}
	if block.Type != "RSA PUBLIC KEY" {
		return nil, errors.New("wrong block type " + block.Type)
	}
	key, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, errors.New("cannot parse public key: " + err.Error())
	}

	return key, nil
}
