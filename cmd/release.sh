#!/bin/sh
env GOOS=linux GOARCH=arm GOARM=7 go build 
scp cmd pi@server:~