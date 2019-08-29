package main

import (
	"flag"
	"fmt"
	"github.com/koykov/mtproto"
	"github.com/koykov/telegram/api"
	"github.com/koykov/telegram/config"
	"log"
	"os"
)

var (
	configFile  = api.FlagStr("config", "c", "", "Path to config file")
	phoneNumber = api.FlagStr("phone", "p", "", "Phone number in international format")
	outputFile  = api.FlagStr("output", "o", "tg.key", "Path to output key file")
)

var (
	conf *config.TGConfig
)

func init() {
	flag.Parse()

	if _, err := os.Stat(*configFile); os.IsNotExist(err) {
		log.Fatal("config file not exists: " + *configFile)
	}
	if _, err := os.Stat(*outputFile); !os.IsNotExist(err) {
		log.Fatal("output file " + *outputFile + " already exists")
	}

	var err error
	conf, err = config.Load(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	conf.RsaKey, err = api.ParseRSAPublicKey(conf.RSAPublicKey)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	mtp, err := mtproto.NewMTProto(conf)
	if err != nil {
		log.Fatal(err)
	}

	err = mtp.Connect()
	if err != nil {
		log.Fatal(err)
	}

	tlCode, err := mtp.AuthSendCode(*phoneNumber)
	if err != nil {
		log.Fatal(err)
	}

	var code int
	fmt.Print("Please enter the login code you received in SMS or Telegram: ")
	_, err = fmt.Scanf("%d", &code)
	if err != nil {
		log.Fatal(err)
	}

	_, err = mtp.AuthSignIn(*phoneNumber, tlCode, code)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Your Telegram key has been saved in %s\n", conf.SecretKey)
}
