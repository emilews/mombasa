package models

import "gopkg.in/mgo.v2/bson"

type Fiat struct {
	FiatTicker	string		`bson:"fiat_ticker"`
	Price 		float64		`bson:"price"`
}

func GetFiatPrice(f string) Fiat{
	result := Fiat{}
	MongoCollection("fiats").Find(bson.M{"fiat_ticker": f}).One(&result)
	return result
}

