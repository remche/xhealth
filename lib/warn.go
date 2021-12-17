package lib

import (
	"context"
	"log"

	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/telegram"
)

func Warn(telegramBot string, receiver int64, message string) {
	telegramService, err := telegram.New(telegramBot)
	if err != nil {
		log.Fatal(err)
	}

	telegramService.AddReceivers(receiver)
	notifier := notify.New()
	notifier.UseServices(telegramService)
	_ = notifier.Send(
		context.Background(),
		"Warning",
		message,
	)

}
