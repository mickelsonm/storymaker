package database

import (
	"os"
	"strings"
	"time"

	"gopkg.in/mgo.v2"
)

var (
	MongoDatabase string
	MongoSession  *mgo.Session

	StoriesCollectionName   = "stories"
	TimelinesCollectionName = "timelines"
)

func InitMongo() error {
	var err error
	if MongoSession == nil {
		connectionString := MongoConnectionString()
		MongoSession, err = mgo.DialWithInfo(connectionString)
		if err == nil {
			MongoDatabase = connectionString.Database
		}
	}
	return err
}

func MongoConnectionString() *mgo.DialInfo {
	var info mgo.DialInfo

	addrs := []string{"127.0.0.1"}
	if hostString := os.Getenv("MONGO_HOSTS"); hostString != "" {
		addrs = strings.Split(hostString, ",")
	}
	info.Addrs = addrs

	info.Username = os.Getenv("MONGO_USERNAME")
	info.Password = os.Getenv("MONGO_PASSWORD")
	info.Database = os.Getenv("MONGO_DATABASE")
	info.Timeout = time.Millisecond * 300
	if info.Database == "" {
		info.Database = "storymaker"
	}

	return &info
}
