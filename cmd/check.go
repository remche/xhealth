package cmd

import (
	"fmt"
	"log"
	"math/big"

	"github.com/remche/xhealth/lib"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:         "check",
	Short:       "One time check",
	Long:        `Runs a one-time check of the xLend health factor`,
	Annotations: map[string]string{"check_args": "yes"},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		rootCmd.PersistentPreRun(cmd, args)
		if (config.TelegramId != 0) != (config.TelegramBotKey != "") {
			log.Fatal("Please provide both telegram-id and telegram-bot-key if you want to enable Telegram notifications")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		health := lib.Query(config)
		fmt.Println(health)
		if config.TelegramId != 0 && health.Cmp(big.NewFloat(float64(config.Treshold))) == -1 {
			lib.Warn(lib.CreateBot(config.TelegramBotKey), config.TelegramId, health.String())
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
