package models

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"log"
	"strconv"
)

type Crypto struct {
	Ticker 			string		`bson:"ticker"`
	LastUpdate 		string 		`bson:"last_update"`
	BCHPrice		float64		`bson:"bch_price"`
	USDPrice		float64		`bson:"usd_price"`
}
type CryptoWithFiat struct {
	Ticker 			string		`bson:"ticker"`
	LastUpdate 		string 		`bson:"last_update"`
	BCHPrice		float64		`bson:"bch_price"`
	USDPrice		float64		`bson:"usd_price"`
	FiatTicker 		string		`bson:"fiat_ticker"`
	FiatPrice 		float64		`bson:"fiat_price"`
}
type CalculatedCrypto struct {
	Ticker 			string		`bson:"ticker"`
	Amount			float64		`bson:"amount"`
	BCHPrice		float64		`bson:"bch_price"`
	USDPrice		float64		`bson:"usd_price"`
}
type CalculatedCryptoWithFiat struct {
	Ticker 			string		`bson:"ticker"`
	Amount			float64		`bson:"amount"`
	BCHPrice		float64		`bson:"bch_price"`
	FiatTicker 		string		`bson:"fiat_ticker"`
	FiatPrice 		float64		`bson:"fiat_price"`
}

func GetCrypto(c string)Crypto {
	result := Crypto{}
	MongoCollection("cryptos").Find(bson.M{"ticker": c}).One(&result)
	fmt.Println(bson.M{"ticker": c})
	return result
}

func GetAllCryptos()[]Crypto{
	results := []Crypto{}
	MongoCollection("cryptos").Find(nil).All(&results)
	return results
}

func GetCryptoWithSpecificFiat(f string, c string) interface{} {
	fiat :=  GetFiatPrice(f)
	crypto := GetCrypto(c)
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

func GetCalculatedCryptoValue(c string, a string)interface{}{
	crypto := GetCrypto(c)
	amount, err := strconv.ParseFloat(a, 64); if err != nil{
		log.Fatal(err)
	}
	calculatedCryptoAmount := amount * crypto.BCHPrice
	calculatedUSDAmount := amount * crypto.USDPrice
	result := CalculatedCrypto{}
	result.Ticker = crypto.Ticker
	result.Amount = amount
	result.BCHPrice = calculatedCryptoAmount
	result.USDPrice = calculatedUSDAmount
	return result
}
