module yagan.bot/cmd

go 1.16

replace yagan.bot/config => ../config

replace yagan.bot/rediscfg => ../rediscfg

replace yagan.bot/telegram => ../telegram

require (
	yagan.bot/config v0.0.0-00010101000000-000000000000
	yagan.bot/mindicator v0.0.0-00010101000000-000000000000 // indirect
	yagan.bot/mysql v0.0.0-00010101000000-000000000000
	yagan.bot/plot v0.0.0-00010101000000-000000000000
	yagan.bot/rediscfg v0.0.0-00010101000000-000000000000
)

replace yagan.bot/mysql => ../mysql

replace yagan.bot/plot => ../plot

replace yagan.bot/mindicator => ../mindicator
