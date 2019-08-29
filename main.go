package main

import (
	"fmt"
	"github.com/koykov/telegram/api"
)

func main() {
	km, err := api.ParseRSAPublicKey("/home/koykov/.ssh/tg0.pub")
	fmt.Println(km, err)
}
