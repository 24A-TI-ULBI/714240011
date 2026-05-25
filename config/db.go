package config

import (
	"gocroot/helper"
	"gocroot/model"
	"os"
)

var MongoString string = getMongoString()

func getMongoString() string {
	val := os.Getenv("MONGOSTRING")
	if val == "" || val == "mongodb://user:pass@host:port" {
		return "mongodb://localhost:27017"
	}
	return val
}

var mongoinfo = model.DBInfo{
	DBString: MongoString,
	DBName:   "iteung",
}

var Mongoconn, _ = helper.MongoConnect(mongoinfo)
