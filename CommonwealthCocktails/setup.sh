#!/bin/bash
#setup you GOPATH here
export GOPATH='/home/ubuntu/workspace/com.roughindustries.cocktails_go/Packages'

#setup the #gitignore filter for lines ending in #gitignore in *.sh, *.go, and *.yaml
git config --global filter.gitignore.clean "sed '/#gitignore$/'d"
git config --global filter.gitignore.smudge cat

go get github.com/mikeflynn/go-alexa
go get github.com/gorilla/mux
go get github.com/codegangsta/negroni
go get github.com/go-sql-driver/mysql
go get github.com/gorilla/context
go get github.com/gorilla/securecookie
go get golang.org/x/oauth2
go get cloud.google.com/go/compute/metadata