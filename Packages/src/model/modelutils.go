// Copyright 2017 Rough Industries LLC. All rights reserved.
//model/modelutils.db.go:package model
package model

import (
	"database/sql"
	"github.com/golang/glog"
)

//Checks the count return for the number of rows
func checkCount(rows *sql.Rows) (count int, err error) {
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			glog.Error(err)
			return 0, err
		}
	}
	return count, nil
}

//Convert a boolean to the corresponding int value, i.e. true = 1 and
//false = 0
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
