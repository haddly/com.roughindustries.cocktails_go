// Copyright 2017 Rough Industries LLC. All rights reserved.
//connectors/db_connection.go: This is a singleton that provides a pool for
//connecting to the database
package connectors_test

import (
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	//_ "github.com/mattn/go-sqlite3"
	"github.com/CommonwealthCocktails/connectors"
	"github.com/spf13/viper"
	"testing"
)

func TestMySQLConnection(t *testing.T) {
	viper.SetConfigName("unit_test_config") // name of config file (without extension)
	viper.AddConfigPath(".")                // optionally look for config in the working directory
	err := viper.ReadInConfig()             // Find and read the config file
	if err != nil {                         // Handle errors reading the config file
		log.Errorf("Fatal error config file: %s \n", err)
		panic(err)
	}
	var dbaddr string
	var dbpasswd string
	var user string
	var proto string
	var port string
	var dbname string
	var dbtype connectors.DBTypesConst
	dbaddr = viper.GetString("dbaddr")
	dbpasswd = viper.GetString("dbpasswd")
	user = viper.GetString("user")
	proto = viper.GetString("proto")
	port = viper.GetString("port")
	dbname = viper.GetString("dbname")
	dbtype = connectors.MySQL
	connectors.SetDBVars(dbaddr, dbpasswd, user, proto, port, dbname, dbtype)
	conn, _ := connectors.GetDBFromMap("www") //get db connection
	if conn == nil {
		t.Errorf("Failed to connect to database")
	} else {
		t.Logf("Successful MySQL DB connection established.")
	}
}
