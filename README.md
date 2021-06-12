# fetch-coin-market-cap-price
`fetch-coin-market-cap-price` is a golang-based repository for getting a real-time cryptocurrency data (current price, marketcap, 24h price, etc) from https://coinmarketcap.com into json.  
`fetch-coin-market-cap-price` also provides a job scheduler allowing you to fetch your data periodically.
## Setup
First, you register for api key in this link https://pro.coinmarketcap.com/

After obtaining the api key and api host, you can fill them in the yaml configuration `config/config.yaml`

```yaml
api_key: "xxx"
api_host: "xxx"
```

## How to run 
```sh
go run main.go
```
