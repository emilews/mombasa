# API 

## Version 1

### GET /api/v1/all
Returns a JSON with an array of all the cryptos in the database.

### GET /api/v1/crypto/
Form data:
```
crypto_ticker -> String
```
Returns a JSON with the requested crypto data.

### GET /api/v1/crypto/amount
Form data:
```
crypto_ticker  -> String
crypto_amount  -> String
```
Returns the requested crypto amount converted to BCH.


### GET /api/v1/fiat/
Form data:
```
crypto_ticker   -> String
fiat_ticker     -> String
```
Returns a JSON with the requested crypto data with its price in the requested fiat.

### GET /api/v1/fiat/amount
Form data:
```
crypto_ticker  -> String
crypto_amount  -> String
fiat_ticker    -> String
```
Returns the requested crypto amount converted to BCH and fiat.
