# fetch-coin-market-cap-price
`fetch-coin-market-cap-price` is a golang-based repository for getting a real-time cryptocurrency data (current price, marketcap, 24h price, etc) from https://coinmarketcap.com into json.  
`fetch-coin-market-cap-price` also provides a job scheduler allowing you to fetch your data periodically. The result ultimately will be stored in `out/coins_market_value.json` as a json format.
## Setup
First, you have to register for api key from this link https://pro.coinmarketcap.com/

After obtaining the api key and api host, you can fill them in the yaml configuration `config/config.yaml`

```yaml
api_key: "xxx"
api_host: "xxx"
```

## How to setup 
```
go mod tidy 
```

## How to run 
```sh
go run main.go
```

## Result
The result json data will be store in `out/coins_market_value.json`
```json
{
    "status":
        {
            "timestamp": "2021-07-14T09:18:51.851Z",
            "error_code":0,
            "error_message":"",
            "elapsed":571,
            "credit_count":25,
            "notice":"",
            "total_count":5671
        },
    "data":
        [
            {
                "id":1,
                "name":"Bitcoin",
                "symbol":"BTC",
                "slug":"bitcoin",
                "num_market_pairs":9021,
                "tags":
                    ["mineable","pow","sha-256","store-of-value","state-channels","coinbase-ventures-portfolio","three-arrows-capital-portfolio","polychain-capital-portfolio","binance-labs-portfolio","arrington-xrp-capital","blockchain-capital-portfolio","boostvc-portfolio","cms-holdings-portfolio","dcg-portfolio","dragonfly-capital-portfolio","electric-capital-portfolio","fabric-ventures-portfolio","framework-ventures","galaxy-digital-portfolio","huobi-capital","alameda-research-portfolio","a16z-portfolio","1confirmation-portfolio","winklevoss-capital","usv-portfolio","placeholder-ventures-portfolio","pantera-capital-portfolio","multicoin-capital-portfolio","paradigm-xzy-screener"],
                    "max_supply":21000000,
                    "circulating_supply":18755925,
                    "total_supply":18755925,
                    "platform":null,
                    "cmc_rank":1,
                    "last_updated":"2021-07-14T09:18:02.000Z",
                    "quote":{
                        "USD":{
                            "price":32481.130921886088,
                            "volume_24h":21102517137.81814,
                            "percent_change_1h":1.76282808,
                            "percent_change_24h":-2.10732618,
                            "percent_change_7d":-6.3076742,
                            "percent_change_30d":-19.02490913,
                            "percent_change_60d":-33.59763771,
                            "percent_change_90d":-48.30690074,
                            "market_cap":609213655486.0763,
                            "last_updated":"2021-07-14T09:18:02.000Z"
                        }
                    }
                        
            }
        ]
}
```
