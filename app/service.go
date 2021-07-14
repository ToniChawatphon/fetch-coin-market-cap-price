package app

import (
	"encoding/json"
	"log"
)

type Service struct {
	api *API
}

// FetchCoinsMarketValue fetches all active cryptocurrencies
// by market cap and return market values in USD
func (s *Service) FetchCoinsMarketValue() {
	log.Println("Fetching data from coin market cap api")
	res, req := s.api.GetMarketPrice(Setting.Url)
	req.Close = true

	defer res.Body.Close()

	s.api.Coins = &CoinMarketCap{}
	err := json.NewDecoder(res.Body).Decode(s.api.Coins)
	if err != nil {
		log.Panicln(res.StatusCode, err.Error())
	}

	data, err := json.Marshal(s.api.Coins)
	if err != nil {
		log.Panicln(err)
	}

	WriteRequestBodyToFile(OutJson, data)
}

func (s *Service) GetCoinMetaData(coin string) {
	log.Printf("Get %v metadata\n", coin)
	for _, c := range *s.api.Coins.Data {
		if coin == c.Symbol {
			log.Printf("- Name: %v\n", c.Name)
			log.Printf("- Symbol: %v\n", c.Symbol)
			log.Printf("- Volume24h: %v\n", c.Quote.USD.Volume24h)
			log.Printf("- PercentChange1h: %v\n", c.Quote.USD.PercentChange1h)
			log.Printf("- PercentChange24h: %v\n", c.Quote.USD.PercentChange24h)
			log.Printf("- PercentChange7d: %v\n", c.Quote.USD.PercentChange7d)
			log.Printf("- PercentChange30d: %v\n", c.Quote.USD.PercentChange30d)
			log.Printf("- PercentChange60d: %v\n", c.Quote.USD.PercentChange60d)
			log.Printf("- PercentChange90d: %v\n", c.Quote.USD.PercentChange90d)
			log.Printf("- MarketCap: %v\n", c.Quote.USD.MarketCap)
			log.Printf("- LastUpdated: %v\n", c.Quote.USD.LastUpdated)
		}
	}
}
