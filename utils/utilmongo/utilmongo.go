package utilmongo

import (
	"log"
	"os"

	mgo "gopkg.in/mgo.v2"
)

var (
	AccessDB string = "access"
)

const (
	MONGO_DB_URL      = "MONGO_DB_URL"
	MONGO_DB_USERNAME = "MONGO_USERNAME"
	MONGO_DB_PASSWORD = "MONGO_PASSWORD"
	MONGO_DB_NAME     = "MONGO_DB_NAME"
)

func MongoDBLogin() (*mgo.Database, error) {
	return loginDB(os.Getenv(MONGO_DB_NAME))
}

func loginDB(databaseName string) (*mgo.Database, error) {
	urlMongo := os.Getenv(MONGO_DB_URL)
	if urlMongo == "" {
		urlMongo = "localhost:27017"
	}
	session, err := mgo.Dial(urlMongo)
	if err != nil {
		return nil, err
	}
	if databaseName == "" {
		databaseName = AccessDB
	}
	if err := session.DB(databaseName).Login(os.Getenv(MONGO_DB_USERNAME), os.Getenv(MONGO_DB_PASSWORD)); err != nil {
		if os.Getenv(MONGO_DB_USERNAME) != "" && os.Getenv(MONGO_DB_PASSWORD) != "" {
			log.Println(err)
		}
	}
	return session.DB(databaseName), nil
}
