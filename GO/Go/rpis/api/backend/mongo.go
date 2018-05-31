package backend

import (
	"rpis/config"
	"labix.org/v2/mgo"
)

func GetSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://"+config.Conf.Mongo.Host)

	if err != nil {
		panic(err)
	}

	return s
}

func GetDB() *mgo.Database {
	session := GetSession()
	return session.DB(config.Conf.Mongo.Name)
}

