package lib

type Config struct {
	CfgFile         string
	Address         string
	Rpc             string
	LpContract      string
	MarketContracts []string
	Treshold        float64
	TelegramBotKey  string
	TelegramId      int64
}
