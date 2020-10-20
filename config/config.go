package config

import (
	"os"
	"log"

	"github.com/joho/godotenv"
	"github.com/namsral/flag"
)

var (
	cfg           = flag.NewFlagSetWithEnvPrefix(os.Args[0], "trader-bot", 0)
	// TelegramBot - bot name
	TelegramBot   *string 
)

func init() {
	err := cfg.Parse(os.Args[1:])
	if err != nil {
		panic(err)
	}

	err = godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	// чтение данных из конфигурационного файла или переменных окружения
	TelegramBot = cfg.String("telegram_bot", os.Getenv("TELEGRAM_BOT"), "Telegram bot name")
  
	// s3Bucket := os.Getenv("S3_BUCKET")
	// secretKey := os.Getenv("SECRET_KEY")
  
	// cfg.VisitAll(func(f *flag.Flag) {
	// 	fmt.Printf("%s : %s\n", f.Name, f.Value)
	// })
}
