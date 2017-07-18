//connectors/mc_connection.go
package connectors

import (
	"github.com/bradfitz/gomemcache/memcache"
	"log"
)

var mc *memcache.Client = nil
var mc_server string

func SetMCVars(in_mc_server string) {
	mc_server = in_mc_server
	log.Println(mc_server)
}

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
