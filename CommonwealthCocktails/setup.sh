#!/bin/bash
export GOPATH='/home/ubuntu/workspace/com.roughindustries.cocktails_go/Packages'
export DBADDR='commonwealthcocktails.c2hmd6oyqsiy.us-east-1.rds.amazonaws.com'
export DBPASSWD='commonwealthcocktails'
export DBUSERNAME='ccuser'
export DBPROTOCOL='tcp'
export DBPORT='3306'
export DBNAME='commonwealthcocktails'
#export  DataSource='Internal'
export DataSource='DB'
go get github.com/mikeflynn/go-alexa
go get github.com/gorilla/mux
go get github.com/codegangsta/negroni
go get github.com/go-sql-driver/mysql