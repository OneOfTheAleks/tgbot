package main

import (
	"flag"
	"fmt"
	"log"
	"tgbot/internal/clients/telegram"
)

const (
	tgUrl = "api.telegram.org"
)

func main() {

	tg := telegram.New(tgUrl, mustToken())
	fmt.Println(tg)
}

func mustToken() string {
	token := flag.String("token", "", " telegram token")
	flag.Parse()
	if *token == "" {
		log.Fatal("token must be")
	}
	return *token
}
