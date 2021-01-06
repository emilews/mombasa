package models

import "gopkg.in/mgo.v2"

func MongoCollection(coll string) *mgo.Collection {
	return MongoStore.C(coll)
}
