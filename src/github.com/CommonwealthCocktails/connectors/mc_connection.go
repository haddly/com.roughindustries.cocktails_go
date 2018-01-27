// Copyright 2017 Rough Industries LLC. All rights reserved.
//connectors/mc_connection.go: This is a singleton that provides a way of
//connecting to the memcache
package connectors

import (
	"github.com/bradfitz/gomemcache/memcache"
	log "github.com/sirupsen/logrus"
)

//memcache variables
var mc *memcache.Client = nil
var mc_server string

//Set the memcache variables for connecting to the memcache server
func SetMCVars(in_mc_server string) {
	mc_server = in_mc_server
	log.Infoln(mc_server)
}

//Get a connection to the memcache
func GetMC() (*memcache.Client, error) {
	if mc_server != "" {
		if mc == nil {
			mc = memcache.New(mc_server)
		}
	} else {
		return nil, nil
	}
	return mc, nil
}
