module yagan.bot/cmd

go 1.16

replace yagan.bot/config => ../config

replace yagan.bot/rediscfg => ../rediscfg

replace yagan.bot/telegram => ../telegram

require (
	yagan.bot/config v0.0.0-00010101000000-000000000000
	yagan.bot/rediscfg v0.0.0-00010101000000-000000000000
)
