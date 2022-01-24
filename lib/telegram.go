package lib

import (
	"log"
	"math/big"
	"time"

	tele "gopkg.in/telebot.v3"
)

func CreateBot(telegramBot string) *tele.Bot {
	pref := tele.Settings{
		Token:  telegramBot,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	return b
}

func StartBot(b *tele.Bot, config *Config) {
	b.Handle("/healthz", func(c tele.Context) error {
		health := Query(config)
		if health.Cmp(big.NewFloat(float64(config.Treshold))) == -1 {
			return c.Send("⚠ XToken health factor\n" + health.String())
		} else {
			return c.Send("ℹ XToken health factor\n" + health.String())
		}
	})

	b.Start()
}

func Warn(b *tele.Bot, receiver int64, message string) {
	user := tele.ChatID(receiver)
	b.Send(user, "⚠ XToken health factor\n "+message, &tele.SendOptions{})

}
