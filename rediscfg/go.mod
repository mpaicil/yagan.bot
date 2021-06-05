module yagan.bot/rediscfg

go 1.16

require (
	github.com/go-redis/redis/v8 v8.8.3
	yagan.bot/config v0.0.0-00010101000000-000000000000
	yagan.bot/telegram v0.0.0-00010101000000-000000000000
)

replace yagan.bot/telegram => ../telegram

replace yagan.bot/config => ../config
