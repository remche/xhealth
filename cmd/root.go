package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/remche/xhealth/lib"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

// var address string
// var rpc string
// var lpContract string
// var marketContract string
// var treshold float64
// var telegramBotKey string
// var telegramId int

var config *lib.Config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "xhealth",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.Address = viper.GetString("address")
		config.Rpc = viper.GetString("rpc")
		config.LpContract = viper.GetString("lp")
		config.MarketContract = viper.GetString("market")
		config.Treshold = viper.GetFloat64("treshold")
		config.TelegramBotKey = viper.GetString("telegram-bot-key")
		config.TelegramId = viper.GetInt64("telegram-id")
		if !common.IsHexAddress(config.Address) {
			log.Fatal("Please provide a valid Ethereum address")
		}
		if !common.IsHexAddress(config.LpContract) {
			log.Fatal("Please provide a valid Ethereum address for liquidity provider contract")
		}
		if !common.IsHexAddress(config.MarketContract) {
			log.Fatal("Please provide a valid Ethereum address for market contract")
		}
		// xor
		if (config.TelegramId != 0) != (config.TelegramBotKey != "") {
			log.Fatal("Please provide bot telegram-id and telegram-bot-key if you want to enable Telegram notifications")
		}
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	config = new(lib.Config)
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.xhealth.yaml)")

	rootCmd.PersistentFlags().StringVarP(&config.Address, "address", "a", "", "address to monitor")
	viper.BindPFlag("address", rootCmd.PersistentFlags().Lookup("address"))

	rootCmd.PersistentFlags().StringVar(&config.Rpc, "rpc", "https://arb1.arbitrum.io/rpc", "rpc endpoint")
	viper.BindPFlag("rpc", rootCmd.PersistentFlags().Lookup("rpc"))

	rootCmd.PersistentFlags().StringVar(&config.LpContract, "lp", "0x8D35b8f4Ee0437EEe49CeA0Dc82F9ba82d52e578", "liquidity provider contract address")
	viper.BindPFlag("lp", rootCmd.PersistentFlags().Lookup("lp"))

	rootCmd.PersistentFlags().StringVar(&config.MarketContract, "market", "0x56F9261EcA26d055A2ca5aa5a6D25A8648C96801", "market contract address")
	viper.BindPFlag("market", rootCmd.PersistentFlags().Lookup("market"))

	rootCmd.PersistentFlags().Float64VarP(&config.Treshold, "treshold", "t", 1.11, "warning treshold")
	viper.BindPFlag("treshold", rootCmd.PersistentFlags().Lookup("treshold"))

	rootCmd.PersistentFlags().StringVar(&config.TelegramBotKey, "telegram-bot-key", "", "telegram bot key")
	viper.BindPFlag("telegram-bot-key", rootCmd.PersistentFlags().Lookup("telegram-bot-key"))

	rootCmd.PersistentFlags().Int64Var(&config.TelegramId, "telegram-id", 0, "telegram id")
	viper.BindPFlag("telegram-id", rootCmd.PersistentFlags().Lookup("telegram-id"))

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".xhealth" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".xhealth")
	}

	viper.SetEnvPrefix("X")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
