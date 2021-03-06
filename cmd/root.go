package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/remche/xhealth/lib"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

var config *lib.Config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "xhealth",
	Short: "Check xLend health factor",
	Long: `xhealth is a CLI application that will let you check and monitor for 
your xLend health factor, optionnaly warning you if a treshold if crossed.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if cmd.Annotations["check_args"] == "yes" {
			config.Address = viper.GetString("address")
			config.Rpc = viper.GetString("rpc")
			config.LpContract = viper.GetString("lp")
			config.MarketContracts = viper.GetStringSlice("market")
			config.Treshold = viper.GetFloat64("treshold")
			config.TelegramBotKey = viper.GetString("telegram-bot-key")
			config.TelegramId = viper.GetInt64("telegram-id")
			if !common.IsHexAddress(config.Address) {
				log.Fatal("Please provide a valid Ethereum address")
			}
			if !common.IsHexAddress(config.LpContract) {
				log.Fatal("Please provide a valid Ethereum address for liquidity provider contract")
			}
			for _, m := range config.MarketContracts {
				if !common.IsHexAddress(m) {
					log.Fatal("Please provide a valid Ethereum address for market contracts")
				}
			}
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

	market := []string{"0x56F9261EcA26d055A2ca5aa5a6D25A8648C96801", // ETH
		"0x769c382124Bd87c78D4316e3dDB77E925c008487", // BTC
		"0xEA32195F8BFb435292aC38659fBB571E6963cFda"} //LINK

	rootCmd.PersistentFlags().StringSliceVar(&config.MarketContracts, "market", market, "market contracts addresses")
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

	viper.SetEnvPrefix("XHEALTH")
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
