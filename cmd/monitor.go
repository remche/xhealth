package cmd

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/remche/xhealth/lib"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var warnDelay int64

var monitorCmd = &cobra.Command{
	Use:         "monitor",
	Short:       "Periodical check",
	Long:        `Monitor xLend health factor`,
	Annotations: map[string]string{"check_args": "yes"},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		rootCmd.PersistentPreRun(cmd, args)
		warnDelay = viper.GetInt64("warn-delay")
		if config.TelegramId == 0 || config.TelegramBotKey == "" {
			log.Fatal("Please provide both telegram-id and telegram-bot-key for monitoring")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		lastWarnTime := time.Date(1982, 1, 15, 0, 0, 0, 0, time.UTC)
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	F:
		for {
			select {
			case <-sigchan:
				fmt.Println("exiting")
				break F
			case <-time.After(10 * time.Second):
				go func() {
					if time.Now().Sub(lastWarnTime).Seconds() >= float64(warnDelay) {
						health := lib.Query(config)
						if health.Cmp(big.NewFloat(float64(config.Treshold))) == -1 {
							lib.Warn(config.TelegramBotKey, config.TelegramId, health.String())
							lastWarnTime = time.Now()
						}
					}
				}()
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(monitorCmd)
	monitorCmd.PersistentFlags().Int64Var(&warnDelay, "warn-delay", 3600, "delay in seconds between warn notifications")
	viper.BindPFlag("warn-delay", monitorCmd.PersistentFlags().Lookup("warn-delay"))
}
