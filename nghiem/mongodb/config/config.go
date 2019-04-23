package config

import "gopkg.in/mgo.v2"

func GetConnection() (*mgo.Database, error) {
	//host := "mongodb://localhost:27017"
	dbName := "session6"
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		return nil, err
	}
	db := session.DB(dbName)
	return db, nil
}
