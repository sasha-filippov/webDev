package main

import (
	"fmt"
	"log"
	"webDevMcLeod/telegram/telegramClient"
)

func main() {

	botToken := "5778811258:AAHw38nJ9hVVvnh1HslfuohwCQfic2UjOZ0"
	//	https://api.telegram.org/bot<token>/METHOD_NAME
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
