package cmd

import (
	"fmt"
	"math/big"

	"github.com/remche/xhealth/lib"
	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		health := lib.Query(config)
		fmt.Println(health)
		if config.TelegramId != 0 && health.Cmp(big.NewFloat(float64(config.Treshold))) == -1 {
			lib.Warn(config.TelegramBotKey, config.TelegramId, health.String())
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
