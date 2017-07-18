#!/bin/bash
#setup you GOPATH here
#SET THIS LINE AND ADD #gitignore to the end of the line as a comment to ignore your info
#export GOPATH=

#setup the #gitignore filter for lines ending in #gitignore in *.sh, *.go, and *.yaml
git config --global filter.gitignore.clean "sed '/#gitignore$/'d"
git config --global filter.gitignore.smudge cat

#get all the libraries we are using
go get github.com/mikeflynn/go-alexa
go get github.com/gorilla/mux
go get github.com/codegangsta/negroni
go get github.com/go-sql-driver/mysql
go get github.com/gorilla/context
go get github.com/gorilla/securecookie
go get github.com/gorilla/sessions
go get golang.org/x/oauth2
go get cloud.google.com/go/compute/metadata
go get github.com/bradfitz/gomemcache/memcache