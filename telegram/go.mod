module yagan.bot/telegram

go 1.16

require (
	github.com/go-telegram-bot-api/telegram-bot-api v4.6.4+incompatible
	github.com/technoweenie/multipartstreamer v1.0.1 // indirect
	yagan.bot/config v0.0.0-00010101000000-000000000000
)

replace yagan.bot/config => ../config
