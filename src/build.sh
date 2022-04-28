#!/bin/sh
#########################################################################
# File Name: build.sh
# Brief: 
# Author:pom
# Created Time: Mon 25 Apr 2022 11:57:16 AM CST
#########################################################################

gofmt -s -w .
go build -o ../bin/server_mcd main.go
