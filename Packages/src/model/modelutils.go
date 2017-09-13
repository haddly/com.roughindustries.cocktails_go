// Copyright 2017 Rough Industries LLC. All rights reserved.
//model/modelutils.db.go:package model
package model

import (
	"database/sql"
	"log"
)

//Checks the count return for the number of rows
func checkCount(rows *sql.Rows) (count int, err error) {
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			log.Fatal(err)
			return 0, err
		}
	}
	return count, nil
}
