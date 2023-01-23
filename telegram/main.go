package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"webDev/telegram/telegramClient"
)

func main() {
	vp := viper.New()
	vp.AddConfigPath("telegram")
	vp.SetConfigType("env")
	vp.SetConfigName("settings")
	err := vp.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	botToken := vp.GetString("Token")

	fmt.Println(botToken)
	////	https://api.telegram.org/bot<token>/METHOD_NAME
	const host = "api.telegram.org"
	offset := 0

	tgClient := telegramClient.New(host, botToken)
	for {
		updates, err := tgClient.Updates(offset)
		if err != nil {
			log.Fatal(err)
		}
		for _, update := range updates {
			err = tgClient.SendMessage(update.Message.Chat.ChatID, update.Message.Text)
			if err != nil {
				fmt.Println(err.Error())
			}
			offset = update.UpdateID + 1
		}
		fmt.Println(updates)
	}
}
