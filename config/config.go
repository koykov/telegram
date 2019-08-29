package config

import (
	"crypto/rsa"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"sync"
)

type TGConfig struct {
	ApiId        int32  `json:"apiId"`
	ApiHash      string `json:"apiHash"`
	RSAPublicKey string `json:"tgPublicKey"`
	SecretKey    string `json:"secretKey"`
	TestAddress  string `json:"testAddress"`
	ProdAddress  string `json:"prodAddress"`

	RsaKey *rsa.PublicKey
}

var (
	conf *TGConfig
	once sync.Once
)

// Singleton get instance implementation.
func Get() *TGConfig {
	once.Do(func() {
		conf = &TGConfig{}
	})
	return conf
}

// Load config from given json file.
func Load(filePath string) (*TGConfig, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, errors.New("file not exists: " + filePath)
	}
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, Get())
	if err != nil {
		return nil, err
	}
	return Get(), nil
}

// Save config to json file.
func Save(filePath string) error {
	bytes, err := json.MarshalIndent(conf, "", "\t")
	if err != nil {
		return err
	}

	if _, err := os.Create(filePath); err != nil {
		return err
	}

	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	if _, err = file.Write(bytes); err != nil {
		return err
	}
	if err = file.Sync(); err != nil {
		return err
	}

	return nil
}
