package main

import (
	"fmt"
	"github.com/Di0niz/trader-bot/config"
)

func main() {
	fmt.Println("hello")
	fmt.Println(*config.TelegramBot)
	fmt.Println("endl")
}