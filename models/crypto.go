package models

import (
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Crypto struct {
	Ticker 			string		`bson:"ticker"`
	LastUpdate 		string 	`bson:"last_update"`
	BCHPrice		float64		`bson:"bch_price"`
	USDPrice		float64		`bson:"usd_price"`
}
type CryptoWithFiat struct {
	Ticker 			string		`bson:"ticker"`
	LastUpdate 		string 	`bson:"last_update"`
	BCHPrice		float64		`bson:"bch_price"`
	USDPrice		float64		`bson:"usd_price"`
	FiatTicker 		string		`bson:"fiat_ticker"`
	FiatPrice 		float64		`bson:"fiat_price"`
}

func GetCrypto(c string)Crypto {
	result := Crypto{}
	MongoCollection("cryptos").Find(bson.M{"ticker": c}).One(&result);
	return result
}

func GetAllCryptos()[]Crypto{
	results := []Crypto{}
	MongoCollection("cryptos").Find(nil).All(&results)
	return results
}

func GetCryptoWithFiat(f string, c string) interface{} {
	fiat :=  GetFiatPrice(f)
	crypto := GetCrypto(c)
	log.Println(crypto)
	transformedPrice := fiat.Price * crypto.USDPrice
	result := CryptoWithFiat{}
	result.Ticker = crypto.Ticker
	result.BCHPrice = crypto.BCHPrice
	result.LastUpdate = crypto.LastUpdate
	result.USDPrice = crypto.USDPrice
	result.FiatTicker = fiat.FiatTicker
	result.FiatPrice = transformedPrice
	return result
}