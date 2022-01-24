# xLend health factor monitor

A small app that allow to monitor your [xToken Lending](https://docs.xtoken.market/products-and-services/xtoken-lending-coming-soon) health factor.

## Usage

```bash
# simple check
xhealth check -a 0xdeadbeef000000000000
# simple check with notification (for cron/job manager use)
xhealth check -a 0xdeadbeef000000000000 --telegram-id 123123 --telegram-bot-key 110201543:AAHdqTcvCH1vGWJxfSeofSAs0K5PALDsaw
# monitor
xhealth monitor -a 0xdeadbeef000000000000 --telegram-id 123123 --telegram-bot-key 110201543:AAHdqTcvCH1vGWJxfSeofSAs0K5PALDsaw
```

```bash
xhealth is a CLI application that will let you check and monitor for 
your XLend health factor, optionnaly warning you if a treshold if crossed.

Usage:
  xhealth [command]

Available Commands:
  check       One time check
  completion  generate the autocompletion script for the specified shell
  help        Help about any command
  monitor     Periodical check

Flags:
  -a, --address string            address to monitor
      --config string             config file (default is $HOME/.xhealth.yaml)
  -h, --help                      help for xhealth
      --lp string                 liquidity provider contract address (default "0x8D35b8f4Ee0437EEe49CeA0Dc82F9ba82d52e578")
      --market string             market contract address (default "0x56F9261EcA26d055A2ca5aa5a6D25A8648C96801")
      --rpc string                rpc endpoint (default "https://arb1.arbitrum.io/rpc")
      --telegram-bot-key string   telegram bot key
      --telegram-id int           telegram id
  -t, --treshold float            warning treshold (default 1.11)

Use "xhealth [command] --help" for more information about a command.
```

## Configuration

Configuration can be set in config file (default to `$HOME/xhealth.yaml`) or environment variables.

```yaml
address: "0xdeadbeef000000000000"
telegram-bot-key: "110201543:AAHdqTcvCH1vGWJxfSeofSAs0K5PALDsaw"
telegram-id: "123123"
treshold: 1.2
```

```bash
export XHEALTH_ADDRESS 0xdeadbeef000000000000
export XHEALTH_TELEGRAM_BOT_KEY 110201543:AAHdqTcvCH1vGWJxfSeofSAs0K5PALDsaw
export XHEALTH_TELEGRAM_ID 123123
```

## How to get you Telegram Bot API Key 

Just send `/newbot` command to [@BotFather](https://t.me/botfather)

## Bot usage

You can query the bot with `/healthz` command.
