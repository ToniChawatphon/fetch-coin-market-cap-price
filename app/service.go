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
