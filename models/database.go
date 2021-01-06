package models

import (
	"gopkg.in/mgo.v2"
	"log"
)

func InitializeMongo() (session *mgo.Session) {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		log.Fatal(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return
}
type MongoStorage struct {
	Session *mgo.Session
}
var MongoStore = MongoStorage{}.Session.DB("mombasa")
