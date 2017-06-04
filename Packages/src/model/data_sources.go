//model/data_sources.go
package model

import ()

var DataSource *DataSourceType = nil

type DataSourceType int

const (
	DB = 1 + iota
	Internal
	YAML
	JSON
)

var DataSourceTypeStrings = [...]string{
	"DB",
	"Internal",
	"YAML",
	"JSON",
}

func (dst DataSourceType) String() string { return DataSourceTypeStrings[dst-1] }

func DSTtoi(dst string) DataSourceType {
	switch dst {
	case "DB":
		return DB
	case "Internal":
		return Internal
	case "YAML":
		return YAML
	case "JSON":
		return JSON
	default:
		return 0
	}
}

func GetDataSourceType() DataSourceType {
	return *DataSource
}

func SetDataSourceType(ds DataSourceType) {
	DataSource = &ds
}
